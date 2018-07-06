class Solution:
    def judgeCircle(self, moves):
        """
        :type moves: str
        :rtype: bool
        """
        pos = [0,0]
        
        for move in moves:
            if move == 'U': pos[0] += 1
            elif move == 'D':   pos[0] -= 1
            elif move == 'L':   pos[1] += 1
            elif move == 'R':   pos[1] -= 1
        
        if pos[0] ==0 and pos[1] == 0:  return True
        else:   return False
