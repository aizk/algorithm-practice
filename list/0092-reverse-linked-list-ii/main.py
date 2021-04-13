# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def reverseBetween(self, head, m, n):
        """
        :type head: ListNode
        :type m: int
        :type n: int
        :rtype: ListNode
        """
        if head is None: return head

        # Move the two pointers until they reach the proper starting point in the list.
        # m 到达指定位置
        cur, prev = head, None
        while m > 1:
            prev = cur
            cur = cur.next
            m, n = m - 1, n - 1

        # The two pointers that will fix the final connections.
        # 一个尾节点
        # 一个 con 表示当前头
        tail, con = cur, prev

        # Iteratively reverse the nodes until n becomes 0.
        while n:
            third = cur.next
            cur.next = prev
            prev = cur
            cur = third
            n -= 1

        # Adjust the final connections as explained in the algorithm
        if con:
            # 此时 prev 作为链表的头部
            con.next = prev
        else:
            # 如果没有头节点，prev 作为头，例如从第一个开始遍历
            head = prev
        # cur 指向后一个节点，使用 tail 连接后面的尾巴
        tail.next = cur
        return head