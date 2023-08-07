package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"strconv"
	"time"
)

func pub(port int, prefix string) {
	ctx, _ := zmq.NewContext()
	defer ctx.Term()

	//PUB 表示publisher角色
	publisher, _ := ctx.NewSocket(zmq.PUB)
	defer publisher.Close()
	//Bind 绑定端口，并指定传输层协议
	publisher.Bind("tcp://127.0.0.1:" + strconv.Itoa(port))

	//publisher会把消息发送给所有subscriber，subscriber可以动态加入
	for i := 0; i < 5; i++ {
		//publisher只能调用send方法
		publisher.Send(prefix+"Hello my followers", 0)
		publisher.Send(prefix+"How are you", 0)
		fmt.Printf("loop %d send over\n", i+1)
		time.Sleep(2 * time.Second)
	}
	publisher.Send(prefix+"END", 0)
}

func sub(port int, prefix string) {
	//SUB 表示subscriber角色
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()

	//Bind 绑定端口，并指定传输层协议
	subscriber.Connect("tcp://127.0.0.1:" + strconv.Itoa(port))
	subscriber.SetSubscribe(prefix) //只接收前缀为prefix的消息
	fmt.Printf("listen to port %d\n", port)

	for {
		//接收广播
		if resp, err := subscriber.Recv(0); err == nil {
			resp = resp[len(prefix):] //去掉前缀
			fmt.Printf("receive [%s]\n", resp)
			if resp == "END" {
				break
			}
		} else {
			fmt.Println(err)
			break
		}
	}
}
