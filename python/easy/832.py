class Solution:
    def flipAndInvertImage(self, A):
        """
        :type A: List[List[int]]
        :rtype: List[List[int]]
        """
        
        return [[ x^1 for x in a][::-1] for a in A]
