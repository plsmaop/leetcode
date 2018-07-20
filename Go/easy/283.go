func moveZeroes(nums []int) {
	last_element := 0

	for index, num := range nums {
		if num != 0 {
			nums[last_element], nums[index] = num, nums[last_element]
			last_element++
		}
	}
}