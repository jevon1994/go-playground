package queue

import "fmt"

type Node struct {
	Next *Node
	Val  interface{}
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) NewNode(val interface{}) *Node {
	return &Node{nil, val}
}
func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

// 尾插
func (q *Queue) Add(val interface{}) {
	node := q.NewNode(val)
	if q.IsEmpty() {
		q.Head = node
		q.Tail = node
		q.Size++
		return
	}
	tail := q.Tail
	tail.Next = node
	q.Tail = node
	q.Size++
}

func (q *Queue) Pop() *Node {
	head := q.Head
	if q.IsEmpty() {
		return head
	}
	next := head.Next
	q.Head = next
	q.Size--
	return head
}

func (q *Queue) Iterate() {
	temp := q.Head
	for {
		if temp == nil {
			break
		}
		fmt.Print(temp.Val, "\n")
		temp = temp.Next
	}
}
