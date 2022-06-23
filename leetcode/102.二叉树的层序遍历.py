#
# @lc app=leetcode.cn id=102 lang=python3
#
# [102] 二叉树的层序遍历
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        ans, level = [], [root] # 声明写在一行
        while root and level: # root 为空在这里判断，不单独写
            ans.append([n.val for n in level]) # 列表推导获取结果
            level = [kid for n in level for kid in (n.left, n.right) if kid] # 列表推导下一层的结果，if kid 过滤 None
        return ans
# @lc code=end

