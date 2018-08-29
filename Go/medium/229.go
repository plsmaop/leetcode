func majorityElement(nums []int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	var candidate1, candidate2, count1, count2 int
	candidate2 = -1
	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1 = 1
		} else if count2 == 0 {
			candidate2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}
	ans := []int{}
	for _, candidate := range []int{candidate1, candidate2} {
		count := 0
		for _, num := range nums {
			if num == candidate {
				count++
			}
		}
		if count > len(nums)/3 {
			ans = append(ans, candidate)
		}
	}
	return ans
}