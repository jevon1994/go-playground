package main

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/streadway/amqp"
	"log"
)

const (
	rabbitMQURL    = "amqp://guest:guest@172.16.255.141:21783/"
	exchangeName   = "example_exchange"
	queueName      = "example_queue"
	bindingKey     = "example_binding_key"
	messageContent = "hello world"
)

func main() {
	err, ch, q := initMQ()
	defer ch.Close()

	// 创建事件总线
	bus := EventBus.New()

	// 注册事件处理函数
	bus.Subscribe("example_event", func(msg string) {
		log.Printf("received message: %v", msg)
	})

	msgs := consume(err, ch, q)

	// 启动事件总线
	startup(msgs, bus)

	// 发布事件
	publish(err, ch)

	// 等待事件处理函数执行完毕
	// 这里只是为了演示，实际使用中不需要等待
	fmt.Scanln()
}

func initMQ() (error, *amqp.Channel, amqp.Queue) {
	// 创建 RabbitMQ 连接
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	// 创建 RabbitMQ 通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a RabbitMQ channel: %v", err)
	}

	// 声明交换机
	err = ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare RabbitMQ exchange: %v", err)
	}

	// 声明队列
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare RabbitMQ queue: %v", err)
	}

	// 绑定队列和交换机
	err = ch.QueueBind(q.Name, bindingKey, exchangeName, false, nil)
	if err != nil {
		log.Fatalf("failed to bind RabbitMQ queue to exchange: %v", err)
	}
	return err, ch, q
}

func publish(err error, ch *amqp.Channel) {
	err = ch.Publish(exchangeName, bindingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(messageContent),
	})
	if err != nil {
		log.Fatalf("failed to publish RabbitMQ message: %v", err)
	}
}

func startup(msgs <-chan amqp.Delivery, bus EventBus.Bus) {
	go func() {
		for delivery := range msgs {
			bus.Publish("example_event", string(delivery.Body))
			delivery.Ack(false)
		}
	}()
}

func consume(err error, ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {
	// 创建消费者
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to register RabbitMQ consumer: %v", err)
	}
	return msgs
}
