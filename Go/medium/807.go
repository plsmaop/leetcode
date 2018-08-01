func maxIncreaseKeepingSkyline(grid [][]int) int {
	var top, left []int
	for i := 0; i < len(grid); i++ {
		max := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		left = append(left, max)
	}
	for i := 0; i < len(grid); i++ {
		max := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		top = append(top, max)
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		right := left[i]
		for j := 0; j < len(grid[i]); j++ {
			min := top[j]
			if min > right {
				min = right
			}
			ans += (min - grid[i][j])
		}
	}
	return ans
}