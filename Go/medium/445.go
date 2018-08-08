/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1_slice, l2_slice []*ListNode
	for l1 != nil {
		l1_slice = append(l1_slice, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		l2_slice = append(l2_slice, l2)
		l2 = l2.Next
	}

	carry, min, max := 0, 0, 0
	var longerSlice *([]*ListNode)
	if len(l1_slice) > len(l2_slice) {
		min = len(l2_slice)
		max = len(l1_slice)
		longerSlice = &l1_slice
	} else {
		max = len(l2_slice)
		min = len(l1_slice)
		longerSlice = &l2_slice
	}

	var ans []int
	for i := 1; i <= min; i++ {
		tmpVal := l1_slice[len(l1_slice)-i].Val + l2_slice[len(l2_slice)-i].Val + carry
		if tmpVal >= 10 {
			ans = append(ans, tmpVal-10)
			carry = 1
		} else {
			ans = append(ans, tmpVal)
			carry = 0
		}
	}
	for i := min + 1; i <= max; i++ {
		tmpVal := (*longerSlice)[len(*longerSlice)-i].Val + carry
		if tmpVal >= 10 {
			ans = append(ans, tmpVal-10)
			carry = 1
		} else {
			ans = append(ans, tmpVal)
			carry = 0
		}
	}
	if carry > 0 {
		ans = append(ans, 1)
	}

	ansRoot := &ListNode{
		Val: ans[len(ans)-1],
	}
	var pointer = ansRoot
	for i := len(ans) - 2; i >= 0; i-- {
		pointer.Next = &ListNode{
			Val: ans[i],
		}
		pointer = pointer.Next
	}

	return ansRoot
}