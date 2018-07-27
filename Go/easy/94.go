/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	pointer := root
	for len(stack) > 0 || pointer != nil {
		for pointer != nil {
			stack = append(stack, pointer)
			pointer = pointer.Left
		}
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, back.Val)
		pointer = back.Right
	}
	return ans
}