package main

import "C"
import (
	"fmt"
)

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
	// ...
}

type Duck struct {
}

func (d *Duck) Quack() {
	fmt.Println("I am quacking! , duck")
}

func (d *Duck) Walk() {
	fmt.Println("I am walking! , duck")
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking! , bird")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking! , bird")
}

var m = make([]interface{}, 0)

func main() {
	b := new(Duck)
	DuckDance(b)
	fmt.Print(len(m))

}
