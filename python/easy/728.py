class Solution:
    def selfDividingNumbers(self, left, right):
        """
        :type left: int
        :type right: int
        :rtype: List[int]
        """
        ans = []
        for x in range(left, right + 1):
            if str(x).count('0') > 0:   continue
            flag = True
            for y in str(x):
                if x % int(y) != 0:
                    flag = False
                    break
            if flag:    ans.append(x)
        return ans
            