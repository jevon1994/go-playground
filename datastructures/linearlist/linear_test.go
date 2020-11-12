package linearlist

import (
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
