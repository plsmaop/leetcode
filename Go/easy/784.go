func letterCasePermutation(S string) []string {
	var ans []string
	var rec func(string, int)
	rec = func(s string, i int) {
		if i == len(S) {
			ans = append(ans, s)
			return
		}
		if s[i] >= 65 && s[i] <= 90 {
			rec(s[:i]+string(s[i]+32)+s[i+1:], i+1)
			rec(s, i+1)
		} else if s[i] >= 97 && s[i] <= 122 {
			rec(s, i+1)
			rec(s[:i]+string(s[i]-32)+s[i+1:], i+1)
		} else {
			rec(s, i+1)
		}
	}
	rec(S, 0)
	return ans
}