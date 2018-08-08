/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(node *TreeNode) bool {
	if node == nil {
		return false
	}
	isLeftValid, isRightValid := dfs(node.Left), dfs(node.Right)
	if !isLeftValid {
		node.Left = nil
	}
	if !isRightValid {
		node.Right = nil
	}

	return node.Val == 1 || isLeftValid || isRightValid
}
func pruneTree(root *TreeNode) *TreeNode {
	dfs(root)
	return root
}