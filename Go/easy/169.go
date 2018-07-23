import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[int(len(nums)/2)]
}