func isMonotonic(A []int) bool {
	if len(A) <= 2 {
		return true
	}
	monotonicType := A[len(A)-1] - A[0]
	formalElm := A[0]
	for i := 1; i < len(A); i++ {
		if A[i]-formalElm > 0 && monotonicType <= 0 {
			return false
		} else if A[i]-formalElm < 0 && monotonicType >= 0 {
			return false
		}
		formalElm = A[i]
	}
	return true
}