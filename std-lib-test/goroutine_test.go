package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	//ch := make(chan int, 100)
	//go send(ch)
	//go recv(ch)
	//go switchAsync(ch)
	//time.Sleep(1e6)
	a := <- future(1)
	fmt.Print(a,time.Now().Nanosecond(),"\n")
	b := <- future(2)
	fmt.Print(b,time.Now().Nanosecond(),"\n")
	tick()
}


func future(a int) chan int{
	fu := make(chan int)
	go func() {
		fu <- a*a
	}()
	return fu
}

func tick(){
	ch := make(chan int, 1)
	go func() { for { ch <- 1 } } ()
L:
	for {
		select {
		case <-ch:
			// do something
		case <-time.After(1):
			// call timed out
			break L
		}
	}
}
func switchAsync(ch chan int) {
	select {
	case u := <-ch:
		fmt.Print(u)
	case k := <-ch:
		fmt.Print(k)
	default:
		fmt.Print("no channel in use")
	}
}

func send(ch chan int) {
	ch <- 1
}

func recv(ch chan int) {
	res := <-ch
	fmt.Print(res)
}
