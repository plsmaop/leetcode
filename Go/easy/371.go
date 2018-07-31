func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	carry, sum := (a&b)<<1, a^b
	return getSum(sum, carry)
}