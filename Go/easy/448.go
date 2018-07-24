func findDisappearedNumbers(nums []int) []int {
	table, ans := make([]bool, len(nums)+1), []int{}
	for _, num := range nums {
		table[num] = true
	}
	for i := 1; i < len(table); i++ {
		if !table[i] {
			ans = append(ans, i)
		}
	}
	return ans
}