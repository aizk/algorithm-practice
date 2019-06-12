package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	result := new(ListNode)
	head := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		node := new(ListNode)
		node.Val = sum % 10
		sum = sum / 10
		result.Next = node
		result = result.Next
	}

	// [5] + [5] = [0, 1]
	if sum ==  1 {
		result.Next = &ListNode{ Val: 1 }
		result = result.Next
	}

	return head.Next
}
