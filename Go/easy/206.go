/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pointer := &head
	table := []*ListNode{*pointer}
	for (*pointer).Next != nil {
		table = append(table, (*pointer).Next)
		pointer = &(*pointer).Next
	}
	for i := len(table) - 1; i > 0; i-- {
		table[i].Next = table[i-1]
	}
	table[0].Next = nil
	return table[len(table)-1]
}