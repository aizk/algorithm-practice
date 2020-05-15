package main

import "fmt"

//根据一棵树的中序遍历与后序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
//
// 返回如下的二叉树：
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// test
	// root 3
	// index 1
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	node := buildTree(inorder, postorder)
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

// 与 105 题类似
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}

	index := 0
	root := postorder[len(postorder) - 1]
	// find root in inorder
	for k, v := range inorder {
		if root == v {
			index = k
			break
		}
	}

	return &TreeNode{
		Val: root,
		Left: buildTree(inorder[:index], postorder[0: index]),
		Right: buildTree(inorder[index + 1:], postorder[index: len(postorder) - 1]),
	}
}
