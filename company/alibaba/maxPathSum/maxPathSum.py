# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def __init__(self):
        self.max_sum = float('-inf')

    def maxPathSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.dfs(root)
        return self.max_sum

    def dfs(self, node):
        if not node:
            return 0
        left = max(0, self.dfs(node.left))
        right = max(0, self.dfs(node.right))
        new_path = node.val + left + right
        self.max_sum = max(self.max_sum, new_path)
        return node.val + max(left, right)