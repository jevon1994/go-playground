package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	n := NewHeap(100)
	//fmt.Print(n)
	n.Insert(89)
	n.Insert(11)
	n.Insert(22)
	n.Insert(55)
	n.Insert(2)
	n.Insert(44)
	n.Insert(1)
	n.Insert(9)
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
