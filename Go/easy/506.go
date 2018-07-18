import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	type indexToScore struct {
		index int
		score int
	}
	table := make([]indexToScore, len(nums))
	for index, score := range nums {
		var tmp indexToScore
		tmp.index = index
		tmp.score = score
		table[index] = tmp
	}
	sort.Slice(table, func(i, j int) bool {
		return table[i].score > table[j].score
	})
	ans := make([]string, len(nums))
	for index, item := range table {
		var content string
		switch index {
		case 0:
			content = "Gold Medal"
		case 1:
			content = "Silver Medal"
		case 2:
			content = "Bronze Medal"
		default:
			content = strconv.Itoa(index + 1)
		}
		ans[item.index] = content
	}
	return ans
}