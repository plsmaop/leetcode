class Solution:
    def check(self, grid, i, j):
        if i < 0 or j < 0 or i > len(grid)-1 or j > len(grid[i])-1:
            return 0
        else:   return grid[i][j]
    def islandPerimeter(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        ans = 0
        
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] != 1:  continue
                ans += (4 - (self.check(grid, i+1, j) +self.check(grid, i-1, j) +self.check(grid, i, j+1) +self.check(grid, i, j-1)))
        return ans
