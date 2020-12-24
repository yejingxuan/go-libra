package libra

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yejingxuan/go-libra/pkg/worker"
	"golang.org/x/sync/errgroup"
	"sync"
)

type Application struct {
	AppName     string          //应用名称
	AppVersion  string          //应用版本号
	wg          *sync.WaitGroup //阻塞锁
	smu         *sync.RWMutex   //互斥锁
	initOnce    sync.Once       //保证init方法只执行一次
	startupOnce sync.Once       //保证start方法只执行一次
	servers     []interface{}   //服务
	workers     []worker.Worker //任务
	//configParser conf.Unmarshaller //配置文件
	HideBanner bool //隐藏Banner
}

func DefaultApplication() *Application {
	app := &Application{}
	app.initialize()
	return app
}

//服务初始化
func (app *Application) initialize() {
	app.initOnce.Do(func() {
		app.smu = &sync.RWMutex{}
		app.workers = make([]worker.Worker, 0)
	})
}

//启动服务
func (app *Application) Startup(fns ...func() error) (err error) {
	//执行系统必须要运行服务
	app.startupOnce.Do(func() {
		err = SerialUntilError(
			//app.parseFlags,
			app.printBanner,
			//app.loadConfig,
			app.startServers,

			/*app.initLogger,
			app.initMaxProcs,
			app.initTracer,
			app.initSentinel,
			app.initGovernor,*/
		)()
	})

	//执行自定义服务
	return SerialUntilError(fns...)()
}

/*func (app *Application) cycleRun(servers func() error) {
	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		servers()
	}()
	app.wg.Wait()
}*/

// 迭代执行
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

//添加server到启动任务
func (app *Application) AppendServers(server ...interface{}) {
	app.servers = append(app.servers, server...)
}

//启动server
func (app *Application) startServers() error {
	var eg errgroup.Group
	for _, item := range app.servers {
		item := item
		switch server := item.(type) {
		case *gin.Engine:
			eg.Go(func() (err error) {
				err = server.Run(fmt.Sprintf(":%d", 8111))
				fmt.Println("success")
				return err
			})
		default:
			fmt.Println("no support")

		}
	}
	return eg.Wait()
}
