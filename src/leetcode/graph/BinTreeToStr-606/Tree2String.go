package BinTreeToStr_606

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return dfs(root)
}

func dfs(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val)
	}
	if root.Right == nil {
		return fmt.Sprintf("%d(%s)", root.Val, dfs(root.Left))
	}
	return fmt.Sprintf("%d(%s)(%s)", root.Val, dfs(root.Left), dfs(root.Right))
}
