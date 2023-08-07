package algorithms

import (
	"fmt"
	"testing"
)

type LRUCache struct {
	Cap   int
	Hash  map[int]*Node
	Cache *DoubleList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:   capacity,
		Hash:  make(map[int]*Node),
		Cache: NewDoubleList(),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Hash[key]
	if !ok {
		return -1
	}
	this.makeRecently(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.Hash[key]
	if ok {
		this.deleteKey(key)
		this.addRecently(key, value)
		return
	}
	if this.Cap == this.Cache.size() {
		this.removeLeastRecently()

	}
	this.addRecently(key, value)
}

func (this *LRUCache) makeRecently(n *Node) {
	this.Cache.remove(n)
	this.Cache.addLast(n)
}

func (this *LRUCache) addRecently(key, val int) {
	n := &Node{
		Key: key,
		Val: val,
	}
	this.Cache.addLast(n)
	this.Hash[key] = n
}

func (this *LRUCache) removeLeastRecently() {
	first := this.Cache.removeFirst()
	delete(this.Hash, first.Key)
}

func (this *LRUCache) deleteKey(key int) {
	node := this.Hash[key]
	this.Cache.remove(node)
	delete(this.Hash, key)

}

type Node struct {
	Key, Val   int
	Next, Prev *Node
}
type DoubleList struct {
	Head, Tail *Node
	Size       int
}

func NewDoubleList() *DoubleList {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Prev = head
	return &DoubleList{
		Head: head,
		Tail: tail,
		Size: 0,
	}
}

func (d *DoubleList) addLast(n *Node) {
	n.Next = d.Tail
	n.Prev = d.Tail.Prev
	d.Tail.Prev.Next = n
	d.Tail.Prev = n
	d.Size++
}

func (d *DoubleList) removeFirst() *Node {
	first := d.Head.Next
	if first == d.Tail {
		return nil
	}
	d.remove(first)
	return first
}

func (d *DoubleList) remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	d.Size--
}

func (d *DoubleList) size() int {
	return d.Size
}

func TestLRU(t *testing.T) {
	//cache1 := Constructor(1)
	//cache1.Put(2, 1)
	//fmt.Println(cache1.Get(2))
	//cache1.Put(3, 2)
	//fmt.Println(cache1.Get(2))
	//fmt.Println(cache1.Get(3))

	constructor := Constructor(2)
	constructor.Put(1, 1)
	constructor.Put(2, 2)
	fmt.Println(constructor.Get(1))
	constructor.Put(3, 3)
	fmt.Println(constructor.Get(2))
	constructor.Put(4, 4)
	fmt.Println(constructor.Get(1))
	fmt.Println(constructor.Get(3))
	fmt.Println(constructor.Get(4))
}
