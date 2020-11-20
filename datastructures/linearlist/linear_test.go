package linearlist

import (
	"go-palyground/datastructures/linearlist/queue"
	"go-palyground/datastructures/linearlist/stack"
	"testing"
)

func TestStack(t *testing.T) {
	stack := stack.NewStack()
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	stack.Push("4")
	stack.Pop()
	stack.Push("5")
	stack.Iterate()
}

func TestQueue(t *testing.T) {
	queue := queue.NewQueue()
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")
	queue.Pop()
}
