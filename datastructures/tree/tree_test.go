package tree

import (
	"fmt"
	"testing"
)

func TestHuffManTree(t *testing.T) {
	heap := NewMinHeap()
	heap.InsertTree(&HuffManTreeNode{Weight: 1})
	heap.InsertTree(&HuffManTreeNode{Weight: 2})
	heap.InsertTree(&HuffManTreeNode{Weight: 3})
	heap.InsertTree(&HuffManTreeNode{Weight: 4})
	heap.InsertTree(&HuffManTreeNode{Weight: 5})
	tree := heap.NewHuffManTree()
	tree.PreTravsel()
}

func TestTree(t *testing.T) {
	tree := NewTree()
	tree.Root = NewTreeNode(4)
	tree.Root.Left = NewTreeNode(2)
	tree.Root.Left.Left = NewTreeNode(1)
	tree.Root.Left.Right = NewTreeNode(3)
	tree.Root.Right = NewTreeNode(6)
	tree.Root.Right.Left = NewTreeNode(5)
	tree.Root.Right.Right = NewTreeNode(7)
	fmt.Print("pre-----", "\n")
	tree.PostOrderTraversal(tree.Root)
	fmt.Print("in-----", "\n")
	tree.PostOrder(tree.Root)
	//fmt.Print("post-----", "\n")
	//tree.PostOrderTraversal(tree.Root)
	//fmt.Print("bfs-----", "\n")
	//tree.BFSTravesal(tree.Root)
	//fmt.Print("dfs-----", "\n")
	//tree.DFSTravesal(tree.Root)
	fmt.Print(tree.FindMax(tree.Root).Val)

}
