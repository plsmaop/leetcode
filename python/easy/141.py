# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        pointer = head
        table = {}
        while True:
            if pointer == None:
                return False
            if table.has_key(pointer):
                return True
            table[pointer] = True
            pointer = pointer.next