package tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewTree()
	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}
	tree.InOrderTraversal(tree.Root)
}
