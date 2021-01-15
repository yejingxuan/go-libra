package main

import (
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/server"
	"github.com/yejingxuan/go-libra/pkg/store/rabbitmq"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	app.AppendServers(diyServer())
	app.Run()
}

//自定义server-接收MQ消息
func diyServer() *server.XServer {
	server, _ := server.StdConfig("").Build(func() {
		log.Info("创建MQ连接")
		conn, err := rabbitmq.StdConfig().Build()
		if err != nil {
			log.Error("创建MQ连接失败", err)
			return
		}
		ch, err := conn.Channel()
		defer ch.Close()
		if err != nil {
			return
		}

		msgs, err := ch.Consume(
			"log-queue-1",
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		forever := make(chan bool)
		go func() {
			for d := range msgs {
				log.Info("接收到消息：" + string(d.Body))
			}
		}()
		<-forever
		return
	})
	return server
}
