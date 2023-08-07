package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var dfs func(root *TreeNode)
	var depth, tmp int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		tmp++
		if root.Left == nil && root.Right == nil {
			depth = max(depth, tmp)
		}
		dfs(root.Left)
		dfs(root.Right)
		tmp--
	}
	dfs(root)
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	t := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	depth := maxDepth(t)
	fmt.Println(depth)
}
