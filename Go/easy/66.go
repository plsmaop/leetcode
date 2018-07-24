func plusOne(digits []int) []int {
	var carry bool
	for i := len(digits) - 1; i >= 0; i-- {
		carry = digits[i]+1 == 10
		if !carry {
			digits[i]++
			break
		}
		digits[i] = 0

	}
	ans := digits
	if carry {
		ans[0] = 0
		ans = append([]int{1}, ans...)
	}
	return ans
}