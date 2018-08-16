/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var cur, prev *TreeNode
	cur = root
	for cur != nil {
		if cur.Right == nil {
			ans = append(ans, cur.Val)
			cur = cur.Left
		} else {
			prev = cur.Right
			for prev.Left != nil && prev.Left != cur {
				prev = prev.Left
			}
			if prev.Left == nil {
				prev.Left = cur
				ans = append(ans, cur.Val)
				cur = cur.Right
			} else {
				prev.Left = nil
				cur = cur.Left
			}
		}

	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}