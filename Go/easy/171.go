func pow(base int, power int) int {
	rtv := 1
	for i := 0; i < power; i++ {
		rtv *= base
	}
	return rtv
}

func titleToNumber(s string) int {
	if len(s) == 1 {
		return int(s[0] - 64)
	}
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		ans += (int(s[i]) - 64) * pow(26, len(s)-1-i)
	}
	return ans
}