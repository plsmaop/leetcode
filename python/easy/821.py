class Solution:
    def shortestToChar(self, S, C):
        """
        :type S: str
        :type C: str
        :rtype: List[int]
        """
        pos = []
        index = 0
        for s in S:
            if s == C:
                pos.append(index)
            index += 1
        ans = []
        for i in range(pos[0]): ans.append(pos[0]-i)
        ans.append(0)
        last_pos = int(pos[0])
        for i in pos[1::]:
            interval = i - last_pos - 1
            end = int(interval/2)
            for y in range(1, end+1):   ans.append(y)
            if interval % 2 == 1:   ans.append(end+1)
            for y in reversed(range(1, 1+end)):   ans.append(y)                
            ans.append(0)
            last_pos = i
        for i in range(1, len(S) - len(ans) + 1):
            ans.append(i)
        return ans