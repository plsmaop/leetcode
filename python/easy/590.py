"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution(object):
    def __init__(self):
        self.ans = []
    def postorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        def dfs(node):
            if node == None:    return
            for child in node.children:
                dfs(child)
            self.ans.append(node.val)
        dfs(root)
        return self.ans