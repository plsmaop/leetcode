class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        ans = 0
        if x < 0:
            ans = int('-' + str(x)[1::][::-1])
        else:   ans = int(str(x)[::-1])
        return ans if  -(1<<31) < ans < (1<<31) - 1 else 0
