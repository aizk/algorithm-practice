# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

def levelOrder(root):
    """
    :type root: TreeNode
    :rtype: List[List[int]]
    """
    l = [root]
    r = []
    while len(l) > 0:
        temp = []
        for node in range(len(l)):
            temp.append(node.val)
            l.pop(0)
            if node.left is not None: l.append(node.left)
            if node.right is not None: l.append(node.right)
        r.append(temp)
    return r


class Solution(object):
    pass


if __name__ == '__main__':
    levelOrder()
