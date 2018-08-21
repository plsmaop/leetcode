/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func find(node *TreeNode, k, ans *int) {
	if node == nil || *k == 0 {
		return
	}
	find(node.Left, k, ans)
	(*k)--
	if *k == 0 {
		*ans = node.Val
	}
	find(node.Right, k, ans)
}

func kthSmallest(root *TreeNode, k int) int {
	ans := 0
	find(root, &k, &ans)
	return ans
}