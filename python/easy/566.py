class Solution:
    def matrixReshape(self, nums, r, c):
        """
        :type nums: List[List[int]]
        :type r: int
        :type c: int
        :rtype: List[List[int]]
        """
        nums_r, nums_c = len(nums), len(nums[0])
        if nums_r*nums_c != r*c:    return nums
        ans = []
        col = []
        for i in nums: 
            for j in i:
                col.append(j)
                if len(col) == c:
                    ans.append(col[:])
                    col = list()
        return ans