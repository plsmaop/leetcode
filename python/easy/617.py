class Solution:
    def mergeTrees(self, t1, t2):
        """
        :type t1: TreeNode
        :type t2: TreeNode
        :rtype: TreeNode
        """
        
        if not t1:  return t2
        if not t2:  return t1
        ans = TreeNode(t1.val + t2. val)
        ans.left = self.mergeTrees(t1.left, t2.left)
        ans.right = self.mergeTrees(t1.right, t2.right)
        return ans
    