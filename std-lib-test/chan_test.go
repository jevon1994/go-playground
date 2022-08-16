package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

var c = make(chan int, 100)

func TestChan(t *testing.T) {

	//x, err := ReadWithSelect(c)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(x)
	i := <-c
	fmt.Println(i)
}

// avoid asleep deadlock
func ReadWithSelect(ch chan int) (x int, err error) {
	timeout := time.NewTimer(time.Microsecond * 500)

	select {
	case x = <-ch:
		return x, nil
	case <-timeout.C:
		return 0, errors.New("read time out")
	}
}
