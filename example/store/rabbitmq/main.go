package main

import (
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/store/rabbitmq"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	app.Run(sendMsg)
	//app.Run(receiveMsg)
}

//发送消息
func sendMsg() error {
	log.Info("创建MQ连接")
	conn, err := rabbitmq.StdConfig().Build()
	if err != nil {
		log.Error("创建MQ连接失败", err)
		return err
	}

	log.Info("创建log-exchange交换机")
	err = rabbitmq.CreateExchange("log-exchange", rabbitmq.KIND_FANOUT, conn)
	if err != nil {
		log.Error("创建log-exchange交换机失败", err)
		return err
	}

	log.Info("创建log-queue-1队列")
	err = rabbitmq.CreateQueueWithEx("log-queue-1", "log-exchange", "", conn)
	if err != nil {
		log.Error("创建log-queue-1队失败", err)
		return err
	}

	log.Info("发送消息")
	err = rabbitmq.SendMsg("log-success222", "log-exchange", "", conn)
	if err != nil {
		log.Error("发送消息失败", err)
		return err
	}
	return err
}

//接收消息
func receiveMsg() error {
	log.Info("创建MQ连接")
	conn, err := rabbitmq.StdConfig().Build()
	if err != nil {
		log.Error("创建MQ连接失败", err)
		return err
	}
	rabbitmq.Receive(func(msg string) {
		log.Info("接收到消息：" + msg)
	}, "log-queue-1", conn)
	return nil
}
