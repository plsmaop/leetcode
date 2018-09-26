func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > cur+nums[i] {
			cur = nums[i]
		} else {
			cur += nums[i]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}