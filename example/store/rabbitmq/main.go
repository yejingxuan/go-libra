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

func receiveMsg() error {
	log.Info("创建MQ连接")
	conn, err := rabbitmq.StdConfig().Build()
	if err != nil {
		log.Error("创建MQ连接失败", err)
		return err
	}
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
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
	return nil
}

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
	err = rabbitmq.SendMsg("log-success", "log-exchange", "", conn)
	if err != nil {
		log.Error("发送消息失败", err)
		return err
	}
	return err
}
