package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
)

func main() {
	// 定义MQTT连接参数
	opts := MQTT.NewClientOptions().AddBroker("tcp://172.16.255.143:31883") // 替换为您的RabbitMQ服务器地址和端口
	opts.SetClientID("mqtt-client")

	// 创建MQTT客户端
	client := MQTT.NewClient(opts)

	// 定义消息接收处理函数
	messageHandler := func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}

	// 连接到MQTT代理服务器
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	topic := "mqtt_topic"
	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Printf("Subscribed to topic: %s\n", topic)

	// 等待程序终止信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// 断开MQTT连接
	//client.Disconnect(250)
	//client.Connect()
}
