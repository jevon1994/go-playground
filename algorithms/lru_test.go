package algorithms

type LRUCache struct {
	Cap   int
	Hash  map[interface{}]*Node
	Cache *DoubleList
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		Cap:   capacity,
		Hash:  make(map[interface{}]*Node),
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
	node, ok := this.Hash[key]
	if ok {
		this.deleteKey(key)
		this.addRecently(key, value)
		return
	}
	if this.Cap == this.Cache.Size {
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
	delete(this.Hash, key)
	this.Cache.remove(node)
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
	d.Tail.Next = n
	d.Tail = n
	d.Size++
}

func (d *DoubleList) removeFirst() *Node {
	if d.Head == d.Tail {
		return nil
	}
	next := d.Head.Next
	d.remove(next)
	return next
}

func (d *DoubleList) remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	d.Size--
}

func (d *DoubleList) size() int {
	return d.Size
}
