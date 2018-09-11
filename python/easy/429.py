"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution(object):
    def levelOrder(self, root):
        """
        :type root: Node
        :rtype: List[List[int]]
        """
        if root == None:    return []
        queue = [root]
        ans = []
        level_num = 1
        while len(queue) > 0:
            tmp_num = level_num
            level_num = 0
            level_elm = []
            while tmp_num > 0:
                top = queue[0]
                level_elm.append(top.val)
                queue = queue[1:]
                for c in top.children: queue.append(c)
                level_num += len(top.children)
                tmp_num -= 1
            ans.append(level_elm)
        return ans
            