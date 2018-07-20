func removeElements(head *ListNode, val int) *ListNode {
	pointer := &head
	for *pointer != nil {
		if (*pointer).Val == val {
			*pointer = (*pointer).Next
		} else {
			pointer = &(*pointer).Next
		}
	}
	return head
}