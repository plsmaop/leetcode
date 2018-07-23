func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lastNum, lastIndex := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != lastNum {
			lastNum = nums[i]
			lastIndex++
			nums[lastIndex] = lastNum
		}
	}
	nums = nums[:lastIndex+1]
	return len(nums)
}