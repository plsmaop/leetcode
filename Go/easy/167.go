func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start+end < 2*len(numbers) && start != end {
		tmp := numbers[start] + numbers[end]
		if tmp == target {
			return []int{start + 1, end + 1}
		} else if tmp > target {
			end--
		} else {
			start++
		}
	}
	return []int{}
}