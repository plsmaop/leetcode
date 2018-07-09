class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        table = {}
        for i, num in enumerate(nums):
            if num not in table:
                table[num] = []
                table[num].append(i)
            else:
                table[num].append(i)
        sort_list = sorted(nums)
        head, tail = 0, len(nums) - 1
        while(True):
            pivot = sort_list[head] + sort_list[tail]
            if pivot == target:
                if sort_list[head] == sort_list[tail]:
                    return [table[sort_list[head]][0], table[sort_list[tail]][1]]
                return [table[sort_list[head]][0], table[sort_list[tail]][0]]
            elif pivot > target: tail -= 1
            else: head += 1
