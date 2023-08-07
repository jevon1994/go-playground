package linearlist

import (
	"testing"
)

func TestMergeList(t *testing.T) {
	node1, node2 := fillListNode(1, 2, 3, 4), fillListNode(2, 3, 4, 5)
	lists := mergeTwoLists(node1, node2)
	printList(lists)
}

func TestPartation(t *testing.T) {
	node := partition(fillListNode(1, 4, 3, 2, 5, 2), 3)
	printList(node)
}
