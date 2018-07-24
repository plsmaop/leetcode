func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var result [256]int
	for i := 0; i < len(s); i++ {
		result[s[i]]++
		result[t[i]]--
	}
	for _, v := range result {
		if v != 0 {
			return false
		}
	}
	return true
}