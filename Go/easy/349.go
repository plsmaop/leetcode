func intersection(nums1 []int, nums2 []int) []int {
	table := make(map[int]bool)
	ans := []int{}
	for _, num := range nums1 {
		table[num] = true
	}
	for _, num := range nums2 {
		if table[num] {
			table[num] = false
			ans = append(ans, num)
		}
	}
	return ans
}   