func missingNumber(nums []int) int {
	sum, largest := 0, 0
	isZeroExisted := false
	for _, num := range nums {
		sum += num
		if num == 0 {
			isZeroExisted = true
		}
		if num > largest {
			largest = num
		}
	}
	ans := (largest+1)*largest/2 - sum
	if !isZeroExisted {
		return 0
	} else if len(nums) == 1 {
		return 1
	}
	if ans <= 0 {
		ans = largest + 1
	}
	return ans
}