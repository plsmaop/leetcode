func minPathSum(grid [][]int) int {
	DP := [][]int{}
	for i, _ := range grid {
		DP = append(DP, []int{})
		for j, _ := range grid[i] {
			top, left := 0, 0
			if i == 0 && j == 0 {
			} else if i == 0 {
				top = DP[i][j-1]
				left = DP[i][j-1]
			} else if j == 0 {
				top = DP[i-1][j]
				left = top
			} else {
				top = DP[i-1][j]
				left = DP[i][j-1]
			}
			top += grid[i][j]
			left += grid[i][j]
			cost := 0
			if top < left {
				cost += top
			} else {
				cost += left
			}
			// fmt.Println(cost)
			DP[i] = append(DP[i], cost)
		}
	}
	return DP[len(grid)-1][len(grid[0])-1]
}