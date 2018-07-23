func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lastNum, lastIndex, isTwo := nums[0], 0, false

	for i := 1; i < len(nums); i++ {
		if !isTwo && lastNum == nums[i] {
			isTwo = true
			lastIndex++
			nums[lastIndex] = lastNum
		}
		if nums[i] != lastNum {
			lastNum = nums[i]
			lastIndex++
			nums[lastIndex] = lastNum
			isTwo = false
		}
	}
	nums = nums[:lastIndex+1]
	return len(nums)
}