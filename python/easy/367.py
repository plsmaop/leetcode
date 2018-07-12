# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def averageOfLevels(self, root):
        """
        :type root: TreeNode
        :rtype: List[float]
        """
        ans, BFS = [], []
        BFS.append(root)
        while len(BFS) > 0:
            tmp, avg = [], 0
            for i in BFS:
                avg += i.val
                if i.left:  tmp.append(i.left)
                if i.right: tmp.append(i.right)
            ans.append(avg/len(BFS))
            BFS = tmp
        return ans