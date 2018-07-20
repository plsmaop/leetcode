func removeElement(nums []int, val int) int {
	index := 0
	for index < len(nums) {
		if nums[index] == val {
			nums[index] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		} else {
			index++
		}
	}
	return len(nums)
}