func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max, min, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max, min = min, max
		}
		if nums[i] > max*nums[i] {
			max = nums[i]
		} else {
			max *= nums[i]
		}
		if nums[i] < min*nums[i] {
			min = nums[i]
		} else {
			min *= nums[i]
		}
		if max > ans {
			ans = max
		}
	}
	return ans
}