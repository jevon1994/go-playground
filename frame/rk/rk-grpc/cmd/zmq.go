package cmd

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"rk-grpc/internal/global"
	"strconv"
)

func InitZmq() {
	//zmqSubscriber := utils.NewZmqSubscriber("tcp://127.0.0.1:5556", "", 100000, time.Second*3, time.Second*3, time.Second*3, time.Second*3, &EventParser{})
	//global.Sub = zmqSubscriber
	//zmqSubscriber.Run()
	sub(5556, "")
}

type EventParser struct{}

func (ep *EventParser) Parse(data []byte) {
	for _, v := range data {
		//event := cloudevents.NewEvent()
		//json.Unmarshal(v, event)
		fmt.Println(v)
	}
}

func sub(port int, prefix string) {
	//SUB 表示subscriber角色
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	//defer subscriber.Close()
	//Bind 绑定端口，并指定传输层协议
	subscriber.Connect("tcp://127.0.0.1:" + strconv.Itoa(port))
	subscriber.SetSubscribe(prefix) //只接收前缀为prefix的消息
	global.Sub = subscriber
	fmt.Printf("listen to port %d\n", port)
}
