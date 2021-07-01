package tree

import (
	"container/list"
	"fmt"
)

var Size int

// BST
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
	node.Right = nil
	node.Left = nil
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

/*
1.叶节点
2.只有一个孩子节点, 父节点指向要删除的孩子节点
3.有左右两颗子树,右子树最小节点或者左子树最大元素代替被删除节点
*/
func (t *Tree) Delete(val int, n *TreeNode) *TreeNode {
	var tmp *TreeNode
	if val < n.Val {
		n.Left = t.Delete(val, n.Left)
	} else if val > n.Val {
		n.Right = t.Delete(val, n.Right)
	} else { // 找到被删除节点
		if n.Right != nil && n.Left != nil { // 左右两个子树
			tmp = t.FindMin(n.Right) // 右子树最小值替代
			n.Val = tmp.Val
			t.Delete(val, n.Right)
		} else { // 只有一个孩子节点或者没有孩子
			if n.Right == nil && n.Left == nil {
				n = nil
			} else if n.Left != nil && n.Right == nil {
				n = n.Left
			} else if n.Left == nil && n.Right != nil {
				n = n.Right
			}
		}
	}
	return n
}

/*----------------------------------- BFS Travesal--------------------------------------*/
// Levelorder
func (t *Tree) BFSTravesal(node *TreeNode) {
	if node == nil {
		return
	}
	q := list.New()
	// 1
	q.PushBack(node)
	for q.Len() != 0 {
		//2
		head := q.Remove(q.Front())
		tempNode := head.(*TreeNode)
		fmt.Print(tempNode.Val, "\n")
		if tempNode.Left != nil {
			//3
			q.PushBack(tempNode.Left)
		}

		if tempNode.Right != nil {
			//4
			q.PushBack(tempNode.Right)
		}
		if q.Len() == 0 {
			fmt.Println()
		}
	}
}

/*
	程序调用通过入栈来实现先调用后返回,递归遍历是系统通过栈实现, 非递归是自己通过栈实现遍历
*/
// recursion
func (t *Tree) InOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	t.InOrderTraversal(node.Left)
	fmt.Print(node.Val, "\n")
	t.InOrderTraversal(node.Right)
}

func (t *Tree) PreOrderTraversal(node *TreeNode) {
	if node != nil {
		fmt.Print(node.Val, "\n")
		t.PreOrderTraversal(node.Left)
		t.PreOrderTraversal(node.Right)
	}
}

func (t *Tree) PostOrderTraversal(node *TreeNode) {
	if node != nil {
		t.PostOrderTraversal(node.Left)
		t.PostOrderTraversal(node.Right)
		fmt.Print(node.Val, "\n")
	}
}

// non-recursion
func (tree *Tree) PreOrder(t *TreeNode) {
	stack := list.New()
	for t != nil || stack.Len() != 0 {
		for t != nil {
			fmt.Print(t.Val, "\n")
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			a := stack.Back()
			t = a.Value.(*TreeNode)
			t = t.Right
			stack.Remove(a)
		}
	}
}

func (tree *Tree) InOrder(t *TreeNode) {
	stack := list.New()
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			a := stack.Back()
			t = a.Value.(*TreeNode)
			fmt.Print(t.Val, "\n")
			t = t.Right
			stack.Remove(a)
		}
	}
}

func (tree *Tree) PostOrder(t *TreeNode) {
	stack := list.New()
	var pre *TreeNode // 第一次访问
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		a := stack.Back()
		top := a.Value.(*TreeNode)
		/*
			   要左右孩子访问之后访问根
			1. 不存在左右孩子, 可以访问
			2. 存在左右孩子, 并且均被访问过一次, 可以访问
			3. 不符合上述条件, 依次右孩子, 左孩子入栈
		*/
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && pre == top.Left) || pre == top.Right {
			fmt.Print(top.Val, "\n")
			pre = top
			stack.Remove(a)
		} else {
			t = top.Right
		}
	}
}

/* DFS Travesal */
func (t *Tree) DFSTravesal(root *TreeNode) {
	if root == nil {
		return
	}
	// 交换左右子树
	tempNode := root.Left
	root.Left = root.Right
	root.Right = tempNode
	t.DFSTravesal(root.Left)
	fmt.Println(root.Val)
	t.DFSTravesal(root.Right)
}

/*--------------------------------------------  Count  ---------------------------------------------------*/
// count
func (t *Tree) CountNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return t.CountNodes(node.Left) + t.CountNodes(node.Right) + 1
}

func (t *Tree) CountLeaves(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Right != nil && node.Left != nil {
		return 1
	}
	return t.CountLeaves(node.Left) + t.CountLeaves(node.Right)
}

func (t *Tree) CountDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := t.CountDepth(node.Left) + 1
	right := t.CountDepth(node.Right) + 1
	if left > right {
		return left
	} else {
		return right
	}
}

func (t *Tree) CountByLevel(node *TreeNode, level int) int {
	if node == nil {
		return 0
	}
	if level == 1 {
		return 1
	}

	return t.CountByLevel(node.Left, level-1) + t.CountByLevel(node.Right, level-1)

}

// find
func (t *Tree) Find(val int) *TreeNode {
	if t == nil {
		return nil
	}
	n := t.Root
	for {
		if val < n.Val {
			n = n.Left
		} else if val > n.Val {
			n = n.Right
		} else if val == n.Val {
			return n
		}
	}

}

func (t *Tree) FindMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (t *Tree) FindMax(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}
