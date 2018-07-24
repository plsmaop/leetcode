func firstUniqChar(s string) int {
	table := make([]int, 26)
	for _, c := range s {
		table[c-97]++
	}
	for index, c := range s {
		if table[c-97] == 1 {
			return index
		}
	}
	return -1
}