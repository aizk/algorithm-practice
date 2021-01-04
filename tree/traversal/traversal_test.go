package traversal

import (
	"testing"
)

var root *TreeNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
	},
	Right: &TreeNode{
		Val: 3,
	},
}

func TestBFS(t *testing.T) {
	result := levelOrder2(root)
	t.Log(result)
}
