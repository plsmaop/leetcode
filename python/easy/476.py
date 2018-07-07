class Solution:
    def findComplement(self, num):
        """
        :type num: int
        :rtype: int
        """
        ans = ''
        for i in '{0:b}'.format(num):
           ans += str(int(i)^1)
        return int(ans, 2)