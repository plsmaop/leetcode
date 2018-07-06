class Solution:
    def peakIndexInMountainArray(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        for index, a in enumerate(A):
            if index == len(A) - 1 or a > A[index + 1]:
                return index
