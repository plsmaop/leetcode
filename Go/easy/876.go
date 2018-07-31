/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	list := []*ListNode{}
	pointer := head
	for pointer != nil {
		list = append(list, pointer)
		pointer = pointer.Next
	}
	return list[len(list)/2]
}