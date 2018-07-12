class Solution:
    def findLUSlength(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: int
        """
        if len(a) != len(b):    return len(a) if len(a) > len(b) else len(b)
        if a == b:  return -1
        return len(a)