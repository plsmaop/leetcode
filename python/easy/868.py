class Solution:
    def binaryGap(self, N):
        """
        :type N: int
        :rtype: int
        """
        S = '{0:b}'.format(N)
        start, ans = -1, 0
        for i in range(len(S)):
            if S[i] == '1':
                if start == -1:
                    start = i
                else:
                    ans = i - start if (i -start) > ans else ans
                    start = i
        return ans
                