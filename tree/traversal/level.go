package traversal


// TODO 是否能不使用 newQueue 变量：
// for range 的过程中修改 slice 会怎样？
// 能够正常结束。循环内改变切片的长度，不影响循环次数，循环次效在循环开始前就已经确定了。
func levelOrder(root *TreeNode) [][]int {
	// 忽略了判断 nil
	if root == nil {
		return nil
	}

	var result [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		var temp []int
		var newQueue []*TreeNode
		for _, node := range queue {
			temp = append(temp, node.Val)
			// 忽略了判断 nil
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		result = append(result, temp)
		queue = newQueue
	}
	return result
}

// 去除 newQueue
func levelOrder2(root *TreeNode) [][]int {
	// 忽略了判断 nil
	if root == nil {
		return nil
	}

	var result [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		var temp []int
		for _, node := range queue {
			temp = append(temp, node.Val)
			// 忽略了判断 nil
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		result = append(result, temp)
		//queue = newQueue
	}
	return result
}