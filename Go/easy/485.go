func findMaxConsecutiveOnes(nums []int) int {
	ans, base := 0, -1
	for index, num := range nums {
		if num == 0 {
			if base != -1 {
				tmp := 0
				if index == base {
					tmp = 1
				} else {
					tmp = index - base
				}
				if ans < tmp {
					ans = tmp
				}
				base = -1
			}
		} else {
			if base == -1 {
				base = index
			}
		}
	}
	if base != -1 {
		tmp := len(nums) - base
		if ans < tmp {
			ans = tmp
		}
	}
	return ans
}