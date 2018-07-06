class Solution:
    def checkPerfectNumber(self, num):
        """
        :type num: int
        :rtype: bool
        """
        if num <= 1: return False
        sum = -num
        for n in range(1, int(math.sqrt(num))+1):
            if num % n == 0:
                sum += (n + num/n)
        return not int(sum)^num