/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	var ans int = 0
	var rec func(*TreeNode) int
	rec = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		right, left := rec(root.Right), rec(root.Left)
		gap := 0
		if left > right {
			gap = left - right
		} else {
			gap = right - left
		}
		ans += gap
		return root.Val + left + right
	}
	rec(root)
	return ans
}