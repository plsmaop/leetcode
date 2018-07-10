class Solution:
    def distributeCandies(self, candies):
        """
        :type candies: List[int]
        :rtype: int
        """
        uniq_num = len(set(candies))
        max_num = int(len(candies) / 2)
        return min(max_num, uniq_num)