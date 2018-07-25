func intersect(nums1 []int, nums2 []int) []int {
	table := make(map[int]int)
	ans := []int{}
	for _, num := range nums1 {
		table[num]++
	}
	for _, num := range nums2 {
		if table[num] > 0 {
			table[num]--
			ans = append(ans, num)
		}
	}
	return ans
}   