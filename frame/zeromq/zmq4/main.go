package main

import (
	"fmt"
	"github.com/pebbe/zmq4"
	"time"
)

type Tick struct {
	Symbol string
	Price  int
}

func (this *Tick) Send(socket *zmq4.Socket) (err error) {
	//fmt.Printf("Send to %s: %q\n", socket, kvmsg.frame)
	_, err = socket.SendMessage(this, 666)
	return
}

func recvtick(socket *zmq4.Socket) (onetick Tick, err error) {
	info, err := socket.RecvMessage(0)
	if err != nil {
		fmt.Println("err in recv", err)
		return
	}
	for _, v := range info {
		fmt.Println(v)

	}
	return
}

func publish() {
	//  Prepare our context and publisher socket
	publisher, _ := zmq4.NewSocket(zmq4.PUB)
	publisher.Bind("tcp://*:5556")

	sequence := int64(1)
	for ; true; sequence++ {
		//  Distribute as key-value message
		tick := Tick{"szse", 200}
		err := tick.Send(publisher)
		fmt.Println("publish", tick)
		if err != nil {
			break
		}
		break
	}
	fmt.Printf("Interrupted\n%d messages out\n", sequence)
}

func subscribe() {
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	subscriber.SetRcvhwm(100000) // or messages between snapshot and next are lost
	subscriber.SetSubscribe("")
	subscriber.Connect("tcp://localhost:5556")

	time.Sleep(time.Second) // or messages between snapshot and next are lost

	//  Now apply pending updates, discard out-of-sequence messages
	for {
		_, err := recvtick(subscriber)
		if err != nil {
			fmt.Println("sub recv err", err)
			break //  Interrupted
		}
	}
}

func main() {
	//fmt.Println("pub")
	//go publish()
	//fmt.Println("sub")
	//subscribe()
	pub(5566, "test-zmq")
	go sub(5566, "test-zmq")
}
