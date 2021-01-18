package rabbitmq

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

const (
	KIND_FANOUT = "fanout"
)

type ConfigRabbitMQ struct {
	Address string
	Timeout int64
}

//标准配置
func StdConfig() ConfigRabbitMQ {
	return rawConfig(fmt.Sprintf("system.rabbitmq"))
}

func rawConfig(name string) ConfigRabbitMQ {
	config := ConfigRabbitMQ{
		Address: viper.GetString(fmt.Sprintf("%s.address", name)),
		Timeout: viper.GetInt64(fmt.Sprintf("%s.timeout", name)),
	}
	return config
}

//创建 mq 连接
func (stdConfig ConfigRabbitMQ) Build() (*amqp.Connection, error) {
	conn, err := amqp.Dial(stdConfig.Address)
	return conn, err
}

//创建exchange交换机
func CreateExchange(exchangeName string, kind string, conn *amqp.Connection) error {
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(
		exchangeName, // name
		kind,         // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	return err
}

//创建队列并与exchange绑定
func CreateQueueWithEx(queueName string, exchangeName string, routingKey string, conn *amqp.Connection) error {
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}
	//声明了队列
	q, err := ch.QueueDeclare(
		queueName, //队列名字为rabbitMQ自动生成
		true,
		false,
		false,
		false,
		nil,
	)
	//交换器跟队列进行绑定，交换器将接收到的消息放进队列中
	err = ch.QueueBind(
		q.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	return nil
}

//发送mq消息
func SendMsg(msg string, exchangeName string, routingKey string, conn *amqp.Connection) error {
	ch, err := conn.Channel()
	defer ch.Close()

	if err != nil {
		return err
	}

	err = ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	return err
}

func Receive(handler func(msg string), queueName string, conn *amqp.Connection) error {
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queueName,
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
			handler(string(d.Body))
		}
	}()
	<-forever
	return nil
}
