class Solution:
    def transpose(self, A):
        """
        :type A: List[List[int]]
        :rtype: List[List[int]]
        """
        ans = []
        for j in A[0]:
            ans.append([])
        for i in range(len(A)):
            for j in range(len(A[i])):
                ans[j].append(A[i][j])
            
        return ans