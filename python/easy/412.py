class Solution:
    def fizzBuzz(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        checker3, checker5 = 0, 0
        ans = []
        for i in range(1, n+1):
            word = ''
            checker3 += 1
            checker5 += 1
            if checker3 == 3:
                checker3  = 0
                word += 'Fizz'
            if checker5 == 5:
                checker5 = 0
                word += 'Buzz'
            if word == '':  word = str(i)
            ans.append(word)
        return ans