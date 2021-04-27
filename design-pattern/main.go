package main

import "fmt"

type Command struct {
	Id1 string
}

type CreatedCommand struct {
	Id string
	Command
}

func get(command interface{}) string {
	return "success"
}

type Event interface {
	handle() int
}

type Event1 struct {
}

func (Event1) handle() int {
	return 1
}

type Event2 struct{}

func (Event2) handle() int {
	return 2
}

func convert(event Event) int {
	return event.handle()
}

func main() {
	c := new(CreatedCommand)
	c.Id = "1"
	fmt.Print(get(c))
	e1 := new(Event1)
	e2 := new(Event1)
	convert(e1)
	convert(e2)
}
