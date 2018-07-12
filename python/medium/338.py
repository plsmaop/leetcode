class Solution:
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        ans = [0]
        digit, carry = 1, 2
        for i in range(1, num+1):
            ans.append(1 + ans[i - digit])
            if i + 1 == carry:
                digit = carry
                carry = (carry<<1)
        return ans 