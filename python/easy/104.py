# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def maxDepth(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        ans = 0
        queue = [root]
        if not root:    return 0
        while len(queue) > 0:
            ans += 1
            tmp = []
            for i in queue:
                if i.left:  tmp.append(i.left)
                if i.right: tmp.append(i.right)
            queue = tmp
        return ans