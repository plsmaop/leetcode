"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution(object):
    def __init__(self):
        self.ans = 0
    def maxDepth(self, root):
        """
        :type root: Node
        :rtype: int
        """
        def dfs(node, depth):
            if node == None:    return
            isLeaf = True
            for children in node.children:
                if children is not None:
                    isLeaf = False
                    dfs(children, depth+1)
            if isLeaf and self.ans < depth+1:
                self.ans = depth+1
        dfs(root, 0)
        return self.ans
            