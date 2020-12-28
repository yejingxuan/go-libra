package libra

import (
	"fmt"
	"net"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/yejingxuan/go-libra/pkg/conf"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/worker"
)

type Application struct {
	AppName     string          //应用名称
	AppVersion  string          //应用版本号
	wg          *sync.WaitGroup //阻塞器
	smu         *sync.RWMutex   //读写锁
	initOnce    sync.Once       //保证init方法只执行一次
	startupOnce sync.Once       //保证start方法只执行一次
	runOnce     sync.Once       //保证run方法只执行一次
	servers     []interface{}   //服务
	workers     []worker.Worker //任务
	confPath    string          //配置文件路径
	logPath     string          //日志路径
	HideBanner  bool            //隐藏Banner
}

func DefaultApplication() *Application {
	app := &Application{}
	app.initialize()
	return app
}

//服务初始化
func (app *Application) initialize() {
	app.initOnce.Do(func() {
		app.wg = &sync.WaitGroup{}
		app.smu = &sync.RWMutex{}
		app.servers = make([]interface{}, 0)
		app.workers = make([]worker.Worker, 0)
	})
}

//准备服务
func (app *Application) Start() (err error) {
	//执行系统必须要运行服务
	app.startupOnce.Do(func() {
		err = SerialUntilError(
			//app.parseFlags,
			app.printBanner,
			app.loadConfig,
			app.initLogger,

			/*app.initLogger,
			app.initMaxProcs,
			app.initTracer,
			app.initSentinel,
			app.initGovernor,*/
		)()
	})
	return nil
}

//启动服务
func (app *Application) Run(fns ...func() error) (err error) {
	//执行系统必须要运行服务
	app.wg.Add(2)
	go func() {
		defer app.wg.Done()
		app.startServers()
	}()
	go func() {
		defer app.wg.Done()
		app.startWorkers()
	}()
	//执行自定义服务
	SerialUntilError(fns...)()
	app.wg.Wait()
	return nil
}

/*func (app *Application) cycleRun(servers func() error) {
	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		servers()
	}()
	app.wg.Wait()
}*/

//迭代执行
func SerialUntilError(fns ...func() error) func() error {
	return func() error {
		for _, fn := range fns {
			if err := fn(); err != nil {
				return err
			}
		}
		return nil
	}
}

//printBanner init
func (app *Application) printBanner() error {
	if app.HideBanner {
		return nil
	}
	const banner = `
 ___        __     _______    _______        __      
|"  |      |" \   |   _  "\  /"      \      /""\     
||  |      ||  |  (. |_)  :)|:        |    /    \    
|:  |      |:  |  |:     \/ |_____/   )   /' /\  \   
 \  |___   |.  |  (|  _  \\  //      /   //  __'  \  
( \_|:  \  /\  |\ |: |_)  :)|:  __   \  /   /  \\  \ 
 \_______)(__\_|_)(_______/ |__|  \___)(___/    \___)
 
 Welcome to libra, starting application ...
`
	fmt.Println(fmt.Sprintf(banner))
	return nil
}

//配置文件初始化
func (app *Application) loadConfig() error {
	err := conf.InitConfig(app.confPath)
	return err
}

//日志初始化
func (app *Application) initLogger() error {
	err := log.InitZapLog()
	return err
}

//添加server到启动任务
func (app *Application) AppendServers(server ...interface{}) error {
	app.smu.Lock()
	defer app.smu.Unlock()
	app.servers = append(app.servers, server...)
	return nil
}

//启动server
func (app *Application) startServers() error {
	var eg errgroup.Group
	for _, item := range app.servers {
		item := item
		switch server := item.(type) {
		case *gin.Engine:
			eg.Go(func() (err error) {
				log.Info(fmt.Sprintf("http-服务加载成功, 访问地址:%d", viper.GetInt("server.http_port")))
				err = server.Run(fmt.Sprintf(":%d", viper.GetInt("server.http_port")))
				return err
			})
		case *grpc.Server:
			eg.Go(func() (err error) {
				log.Info(fmt.Sprintf("grpc-服务加载成功, 访问地址:%s, 协议:%s", viper.GetString("server.grpc_port"), viper.GetString("server.grpc_protocol")))
				listen, err := net.Listen(viper.GetString("server.grpc_protocol"), fmt.Sprintf(":%d", viper.GetInt("server.grpc_port")))
				if err != nil {
					log.Error("listen err:", err)
					return err
				}
				err = server.Serve(listen)
				return err
			})
		default:
			log.Info("no support")
		}
	}
	return eg.Wait()
}

//添加server到启动任务
func (app *Application) AppendWorkes(worker ...worker.Worker) error {
	app.smu.Lock()
	defer app.smu.Unlock()
	app.workers = append(app.workers, worker...)
	return nil
}

//启动workers
func (app *Application) startWorkers() error {
	var eg errgroup.Group
	for _, item := range app.workers {
		worker := item
		eg.Go(func() (err error) {
			c := cron.New()
			c.AddFunc(worker.Config.WorkerCron, worker.GetTask())
			c.Start()
			return nil
		})
	}
	return eg.Wait()
}

func (app *Application) SetConfigPath(confPath string) error {
	if confPath != "" {
		app.confPath = confPath
	}
	return nil
}

func (app *Application) SetLogPath(logPath string) error {
	if logPath != "" {
		app.logPath = logPath
	}
	return nil
}
