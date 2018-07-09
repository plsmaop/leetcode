class Solution:
    def calPoints(self, ops):
        """
        :type ops: List[str]
        :rtype: int
        """
        last = 0
        second_last = 0
        score = []
        for op in ops:
            if op == 'C':
                score.pop()
            elif op == 'D':
                score.append(score[len(score) - 1]*2)
            elif op == '+':
                score.append(score[len(score) - 1] + score[len(score) - 2])
            else:
                score.append(int(op))
                
        return sum(score)
