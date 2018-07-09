class Solution:
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        ans = nums[0]
        for i in range(1, len(nums)):
            ans^=nums[i]
        return ans

# faster one

class Solution:
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        table = {}
        for num in nums:
            if num not in table:
                table[num] = 1
            else:   table[num] += 1
        for num in nums:
            if table[num] == 1: return num
