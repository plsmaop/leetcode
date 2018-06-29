class Solution:
    def findPeakElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        
        for index, num in enumerate(nums):
            if len(nums) == 1:  return index
            elif index == 0 and num > nums[index + 1]: return index
            elif index == len(nums) - 1 and num > nums[index - 1]:  return index
            elif num > nums[index + 1] and num > nums[index - 1]: return index
