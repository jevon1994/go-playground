package main

import (
	"fmt"
	"testing"
)

//1
type Template interface {
	do()
	exec()
}

//2
type Abstract struct {
}

func (a *Abstract) do() {
	fmt.Println("do")
}

func (a *Abstract) exec() {
	fmt.Println("exec")
}

//3
type Concrete struct {
	Abstract
}

//override
func (a *Concrete) do() {
	fmt.Println("Just Do IT")
}

//4
type UseTemplate struct {
	Template
}

func (t *UseTemplate) DoAndExec() {
	t.do()
	t.exec()
}

func TestTemplate(t *testing.T) {
	c := new(Concrete)
	//c.do()
	u := UseTemplate{c}
	u.DoAndExec()
}
