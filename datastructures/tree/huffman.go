package tree

import (
	"fmt"
)

type HuffManTreeNode struct {
	Left, Right *HuffManTreeNode
	Weight      int
}
type MinHeap struct {
	Size        int
	HuffManTree []*HuffManTreeNode // 存储哈夫曼树
}

func NewMinHeap() *MinHeap {
	tree := make([]*HuffManTreeNode, 1)
	//tree = append(tree, &HuffManTreeNode{})
	tree[0] = &HuffManTreeNode{}
	return &MinHeap{
		Size:        0,
		HuffManTree: tree,
	}
}

//O(NLogN)
func (m *MinHeap) InsertTree(tree *HuffManTreeNode) {
	m.Size++
	i := m.Size
	m.HuffManTree = append(m.HuffManTree, &HuffManTreeNode{})
	//向上找比 tree 大的
	for ; m.HuffManTree[i/2].Weight > tree.Weight; i /= 2 {
		m.HuffManTree[i] = m.HuffManTree[i/2] // 找到比 tree 大的放到下面
	}
	m.HuffManTree[i] = tree // 找到 data[i/2] 比 val 小的 i
}

func (m *MinHeap) DeleteMin() *HuffManTreeNode {
	var parent, child int
	Min := m.HuffManTree[1]
	tmp := m.HuffManTree[m.Size].Weight
	m.Size -= 1
	//是否有左子树
	for parent = 1; parent*2 <= m.Size; parent = child {
		child = parent * 2
		//child 不为当前最后一个结点，即 parent 有右孩子结点
		if child != m.Size && m.HuffManTree[child].Weight > m.HuffManTree[child+1].Weight {
			child += 1 // 返回左右孩子较小的
		}
		// 最后元素都比左右孩子大, 不需要调整
		if tmp <= m.HuffManTree[child].Weight {
			break
		} else {
			//否则把小的孩子提回来
			m.HuffManTree[parent] = m.HuffManTree[child]
		}
	}
	m.HuffManTree[parent].Weight = tmp
	fmt.Print("delete min : ", Min.Weight, "\n")
	return Min
}

func (m *MinHeap) NewHuffManTree() *HuffManTreeNode {
	for m.Size > 1 {
		node := &HuffManTreeNode{}
		node.Left = m.DeleteMin()
		node.Right = m.DeleteMin()
		node.Weight = node.Left.Weight + node.Right.Weight
		m.InsertTree(node)
	}
	return m.DeleteMin()
}

func (m *HuffManTreeNode) PreTravsel() {
	if m != nil {
		fmt.Print(m.Weight, "\n")
		m.Left.PreTravsel()
		m.Right.PreTravsel()
	}
}

func (m *HuffManTreeNode) InOrderTravsel() {
	if m != nil {
		m.Left.PreTravsel()
		fmt.Print(m.Weight, "\n")
		m.Right.PreTravsel()
	}
}
