func containsDuplicate(nums []int) bool {
	table := make(map[int]bool)
	for _, num := range nums {
		if table[num] {
			return true
		}
		table[num] = true
	}
	return false
}