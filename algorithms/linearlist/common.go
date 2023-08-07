package linearlist

import "fmt"

func printList(node *ListNode) {
	if node == nil {
		return
	}
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func fillListNode(nums ...int) *ListNode {
	l := &ListNode{}
	p := l
	for _, num := range nums {
		p.Next = &ListNode{Val: num}
		p = p.Next
	}
	return l.Next
}
