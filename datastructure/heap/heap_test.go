package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	n := NewHeap(100)
	//fmt.Print(n)
	n.InsertWithUp(89)
	n.InsertWithUp(88)
	n.Insert(87)
	n.BFSTravesal()
	//n.InsertWithUp(86)
	//n.BFSTravesal()
	//n.Insert(85)
	//n.Insert(84)
	//n.InsertWithUp(83)
	//n.Insert(82)
	//n.BFSTravesal()
	//n.BFSTravesal()
	//n.DeleteMax()
	//n.BFSTravesal()
	//n.DeleteMax()
	//n.BFSTravesal()
	//n.DeleteMax()
	//n.BFSTravesal()
	//n.DeleteMax()
	//n.BFSTravesal()
}

//n 个 元素插入  O(N * LogN)
func MaxHeapWithLogN(heap []int) {
	cap := len(heap)
	h := NewHeap(cap)
	for i := 0; i < cap; i++ {
		h.Insert(heap[i])
	}
}
