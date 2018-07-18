import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	var preOrder func(*TreeNode, string)
	preOrder = func(root *TreeNode, s string) {
		if root == nil {
			return
		}

		s += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			ans = append(ans, s)
			return
		}
		preOrder(root.Left, s+"->")
		preOrder(root.Right, s+"->")
	}

	preOrder(root, "")
	return ans
}