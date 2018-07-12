class Solution:
    def countPrimeSetBits(self, L, R):
        """
        :type L: int
        :type R: int
        :rtype: int
        """
        ans = 0
        table = dict()
        for i in range(L, R+1):
            n = '{0:b}'.format(i).count('1')
            flag = True if n != 1 else False
            if n not in table:
                for j in range(2, int(math.sqrt(n)) + 1):
                    if n % j == 0:
                        flag = False
                        break
                if flag:
                    ans += 1
                    table[n] = True
            else:   ans += 1
        return ans