# Definition for a binary tree node.
import collections


# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def levelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []
        res = []
        q = collections.deque()
        q.append(root)
        while len(q) > 0:
            level = []
            tmp = []
            while len(q) > 0:
                node = q.popleft()
                level.append(node.val)
                if node.left:
                    tmp.append(node.left)
                if node.right:
                    tmp.append(node.right)
            q.extend(tmp)
            res.append(level)
        return res