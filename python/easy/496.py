class Solution:
    def nextGreaterElement(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        table = dict()
        ans = []
        for index, num in enumerate(nums2):
            if num not in table:    table[num] = index
        for num in nums1:
            tmp = -1
            next_index = table[num] + 1
            nums2_len = len(nums2)
            while next_index < nums2_len:
                if nums2[next_index] > num:
                    tmp = nums2[next_index]
                    break
                next_index += 1
            ans.append(tmp)
        return ans