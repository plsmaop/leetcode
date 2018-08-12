/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		top := queue[0]
		if top.Val == val {
			return top
		}
		queue = queue[1:]
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return nil
}