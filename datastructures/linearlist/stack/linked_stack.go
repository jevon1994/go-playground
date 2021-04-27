package stack

import "fmt"

// 链式实现
type Node struct {
	Next *Node
	Val  interface{}
}

type Stack struct {
	Size int
	Top  *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) Push(val interface{}) {
	n := new(Node)
	n.Val = val
	n.Next = s.Top
	s.Top = n
	s.Size++
}

func (s *Stack) Pop() *Node {
	top := s.Top
	if top == nil {
		return top
	}
	next := top.Next
	s.Top = next
	s.Size--
	return top
}

func (s *Stack) GetTop() *Node {
	return s.Top
}

func (s *Stack) Iterate() {
	temp := s.Top
	for {
		if temp.Next == nil {
			fmt.Print(temp.Val, "\n")
			break
		}
		fmt.Print(temp.Val, "\n")
		temp = temp.Next
	}
}
