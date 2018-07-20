# Note
## 203
```
Remove all elements from a linked list of integers that have value val.
```

my sol:
```
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
```
* The variable, `pointer`, is a pointer which points to the pointer of `ListNode`, so `*pointer` is the pointer of `ListNode`.
* Once `val` is equal to current `(*pointer).Val`, the code will change the value of `(*pointer)`, making  `(*pointer)` point to `(*pointer).Next`. That is, the current `ListNode` `*pointer` used to point is no longer being pointed
.
* Otherwise, the pointer of pointer, `pointer`, will point to the pointer of the next node of `*pointer`
