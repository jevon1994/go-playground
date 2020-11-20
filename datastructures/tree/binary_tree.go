package tree

import "fmt"

var Size int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func NewTree() *Tree {
	return &Tree{}
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func (t *Tree) IsEmpty() bool {
	return Size == 0
}

// BST
func (t *Tree) Insert(val int) {
	node := NewTreeNode(val)
	node.Val = val
	if t.Root == nil {
		t.Root = node
	} else {
		temp := t.Root
		for temp != nil {
			if val < temp.Val {
				if temp.Left == nil {
					temp.Left = node
					return
				} else {
					temp = temp.Left
				}
			} else {
				if temp.Right == nil {
					temp.Right = node
					return
				} else {
					temp = temp.Right
				}
			}
		}
	}
}

func (t *Tree) InOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, "\n")
	t.InOrderTraversal(node.Left)
	t.InOrderTraversal(node.Right)
}
