func sortArrayByParity(A []int) []int {
	var even, odd []int
	for _, a := range A {
		if a&1 == 1 {
			odd = append(odd, a)
		} else {
			even = append(even, a)
		}
	}
	return append(even, odd...)
}