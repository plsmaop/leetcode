/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	carry, ans := 0, []int{}
	for p1 != nil || p2 != nil {
		currentDigit := carry
		if p1 != nil {
			currentDigit += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			currentDigit += p2.Val
			p2 = p2.Next
		}
		if currentDigit-10 >= 0 {
			carry = 1
			ans = append(ans, currentDigit-10)
		} else {
			carry = 0
			ans = append(ans, currentDigit)
		}
	}
	if carry == 1 {
		ans = append(ans, 1)
	}
	ansNode := ListNode{
		Val: ans[0],
	}
	pointer := &ansNode
	for i := 1; i < len(ans); i++ {
		newNode := ListNode{
			Val: ans[i],
		}
		pointer.Next = &newNode
		pointer = &newNode
	}
	return &ansNode
}