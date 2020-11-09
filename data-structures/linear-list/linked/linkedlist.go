package linked

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

type Linkedlist struct {
	size int
	head *Node
}

func EmptyLinkedlist() *Linkedlist {
	return &Linkedlist{}
}

func EmptyNode(val interface{}) *Node {
	return &Node{val, nil}
}

func (l *Linkedlist) HeadAdd(val interface{}) {
	node := EmptyNode(val)
	node.next = l.head
	l.head = node
	l.size++
}

func (l *Linkedlist) TailAdd(val interface{}) {
	node := EmptyNode(val)
	if l.head == nil {
		l.head = node
		l.size++
		return
	}
	cur := l.head
	for ; cur.next != nil; cur = cur.next {

	}
	cur.next = node
	l.size++
}

func (l *Linkedlist) HeadDel() interface{} {
	if l.head == nil {
		return l.head
	}
	tem := l.head
	l.head = tem.next
	l.size--
	return tem.val
}

func (l *Linkedlist) tailDel() interface{} {
	// 检查头结点
	if l.head == nil {
		return l.head.val
	}
	// 是否为空表
	if l.head.next == nil {
		return l.HeadDel()
	}
	// 从尾部删除元素
	tem := l.head
	for ; tem.next.next != nil; tem = tem.next {

	}
	val := tem.next.val
	tem.next = nil
	l.size--
	return val
}
func (l *Linkedlist) Reverse() {
	var pre, rNext *Node
	cur := l.head
	for cur != nil {
		// 自定义 next
		rNext = cur.next
		// 自定义 头节点
		cur.next = pre
		// 更换头结点
		pre = cur
		cur = rNext
	}
	l.head = pre
}

func (l *Linkedlist) PrintOut() {
	for a := l.head; a != nil; a = a.next {
		fmt.Printf("%s \n", a.val)
	}
}
