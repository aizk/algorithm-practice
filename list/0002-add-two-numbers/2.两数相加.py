#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] 两数相加
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        dummy = tail = ListNode(0)
        carry = 0
        while l1 or l2 or carry:
            carry += (l1.val if l1 else 0) + (l2.val if l2 else 0)
            dummy.next = ListNode(carry % 10)
            dummy = dummy.next
            carry = carry // 10 # 是否进位
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        return tail.next
# @lc code=end

