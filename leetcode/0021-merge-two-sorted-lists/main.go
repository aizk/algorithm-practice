package main

import "log"

func main() {
	//l1 := &ListNode{Val: 1}
	//l1.Next = &ListNode{Val: 2}
	//l1.Next.Next = &ListNode{Val: 4}
	//
	//l2 := &ListNode{Val: 1}
	//l2.Next = &ListNode{Val: 3}
	//l2.Next.Next = &ListNode{Val: 4}

	l1 := &ListNode{Val: -9}
	l1.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 7}

	//for l1 != nil {
	//	log.Println(l1.Val)
	//	l1 = l1.Next
	//}

	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		log.Println(l3.Val)
		l3 = l3.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
// 需要判断多个边界条件
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	start := new(ListNode)
	next := start
	for l1 != nil || l2 != nil {
		if l1 == nil {
			next.Val = l2.Val
			l2 = l2.Next
			if l2 == nil {
				return start
			} else {
				next.Next = new(ListNode)
				next = next.Next
			}
			continue
		}

		if l2 == nil {
			next.Val = l1.Val
			l1 = l1.Next
			if l1 == nil {
				return start
			} else {
				next.Next = new(ListNode)
				next = next.Next
			}
			continue
		}

		if l1.Val <= l2.Val {
			next.Val = l1.Val
			l1 = l1.Next
		} else {
			next.Val = l2.Val
			l2 = l2.Next
		}
		next.Next = new(ListNode)
		next = next.Next
	}
	return start
}

// 递归的实现
// 假设函数可以排序，当 l1 小于 l2 时，l1 指向 递归函数返回的 l1 next 和 l2 的结果，反之，最终能得到结果
func mergeTwoListsRecurse(l1 *ListNode, l2 *ListNode) *ListNode {
	// 都不为 nil 的时候简化问题
	if l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l1.Next = mergeTwoListsRecurse(l1.Next, l2)
			return l1
		} else {
			l2.Next = mergeTwoListsRecurse(l1, l2.Next)
			return l2
		}
	}
	// 退出条件
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	return nil
}