package traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) []int {
	var res []int
	return res
}

func helper(root *TreeNode, res *[]int) {
	*res = append(*res, root.Val)
}