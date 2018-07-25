/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	node := []*ListNode{head}
	pointer := &head
	for (*pointer).Next != nil {
		pointer = &(*pointer).Next
		node = append(node, *pointer)
	}
	if len(node)-n-1 < 0 {
		return node[0].Next
	}
	node[len(node)-n-1].Next = node[len(node)-n-1].Next.Next
	return head
}