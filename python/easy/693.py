class Solution:
    def hasAlternatingBits(self, n):
        """
        :type n: int
        :rtype: bool
        """
        bit = '{0:b}'.format(n)
        for i in range(1, len(bit)):
            if int(bit[i-1])^int(bit[i]) == 0:  return False
        return True