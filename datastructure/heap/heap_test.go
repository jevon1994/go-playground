package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	n := NewHeap(100)
	//fmt.Print(n)
	n.Insert(89)
	n.Insert(88)
	n.Insert(87)
	n.Insert(86)
	n.Insert(85)
	n.Insert(84)
	n.Insert(83)
	n.Insert(82)
	n.BFSTravesal()
	n.DeleteMax()
	n.BFSTravesal()
	n.DeleteMax()
	n.BFSTravesal()
	n.DeleteMax()
	n.BFSTravesal()
	n.DeleteMax()
	n.BFSTravesal()
}

//n 个 元素插入  O(N * LogN)
func MaxHeapWithLogN(heap []int) {
	cap := len(heap)
	h := NewHeap(cap)
	for i := 0; i < cap; i++ {
		h.Insert(heap[i])
	}
}
