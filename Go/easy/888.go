import "sort"

func fairCandySwap(A []int, B []int) []int {
	sort.Ints(A)
	sort.Ints(B)
	var sumA, sumB int
	for _, a := range A {
		sumA += a
	}
	for _, b := range B {
		sumB += b
	}
	index1, index2 := 0, 0
	for index1 < len(A) && index2 < len(B) {
		if sumA-2*A[index1] == sumB-2*B[index2] {
			break
		} else if sumA-2*A[index1] > sumB-2*B[index2] {
			index1++
		} else {
			index2++
		}
	}
	return []int{A[index1], B[index2]}
}