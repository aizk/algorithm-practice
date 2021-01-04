package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sums []int

func main() {
	//n := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//	},
	//}

	n1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{Val: 5},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	//fmt.Println(sumNumbers(n))
	fmt.Println(sumNumbers(n1))
}

func sumNumbers(root *TreeNode) int {
	sums = []int{}
	if root == nil {
		return 0
	}
	helper(root, &[]int{})
	s := 0
	for _, sum := range sums {
		s += sum
	}

	return s
}

// 前序遍历
func helper(node *TreeNode, path *[]int) {
	// 左边 over
	if node.Left == nil && node.Right == nil {
		// record 路径
		sums = append(sums, getPathNum(*path, node.Val))
		return
	}

	if node.Left != nil {
		*path = append(*path, node.Val)
		helper(node.Left, path)
		*path = (*path)[:len(*path) - 1]
	}
	if node.Right != nil {
		*path = append(*path, node.Val)
		helper(node.Right, path)
		*path = (*path)[:len(*path) - 1]
	}
}

func getPathNum(path []int, x int) int {
	val := 0
	for _, v := range path {
		val = val * 10 + v
	}
	val = val * 10 + x
	return val
}