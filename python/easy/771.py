class Solution:
    def numJewelsInStones(self, J, S):
        """
        :type J: str
        :type S: str
        :rtype: int
        """
        table = [False] * 70 
        number_of_j = 0
        for j in J:
            table[ord(j) - 65] = True 
        for s in S:
            if table[ord(s) - 65]:
                number_of_j += 1
        return number_of_j
            
        
