# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        dumy = None
        dumy.next = head
        pre = dumy
        while head:
            # 归并到翻转部分了..
            tail = pre
            for i in range(k):
                tail = tail.next
                if not tail:
                    return dumy.next

            nex = tail.next
            head, tail = self.reverse(head, tail)
            # 把子链表重新接回原链表
            pre.next = head
            tail.next = nex

            # new round
            pre = tail
            head = tail.next

        return dumy.next

    # 部分翻转
    def reverse(self, head, tail):
        # 末尾节点作为头节点
        prev = tail.next
        cur = head
        while prev != tail:
            nex = cur.next
            cur.next = prev
            prev = cur
            cur = nex
        return tail, head