package main

//根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//
// 返回如下的二叉树：
//
//     3
//   / \
//  9  20
//    /  \
//   15   7

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	// test
	// root 3
	// index 1
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	node := buildTree(preorder, inorder)
	ensure(node, 3)
	ensure(node.Left, 9)
	ensure(node.Right, 20)
	ensure(node.Right.Left, 15)
	ensure(node.Right.Right, 7)
}

func ensure(node *TreeNode, value int)  {
	fmt.Println("check node value", node.Val, value)
	if node.Val != value {
		fmt.Println("error node value", value)
	}
}

// 关键就是 preorder 的第一个元素是 root 节点
// 然后到 inorder 里面划分数组，子问题就出来了
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	index := 0
	root := preorder[0]
	// find root in inorder
	for k, v := range inorder {
		if root == v {
			index = k
			break
		}
	}

	return &TreeNode{
		Val: preorder[0],
		Left: buildTree(preorder[1: 1 + index], inorder[:index]),
		Right: buildTree(preorder[1 + index:], inorder[index + 1:]),
	}
}