/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		top1, top2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		if top1 != nil && top2 != nil {
			if top1.Val != top2.Val {
				return false
			}
			queue1 = append(queue1, top1.Left, top1.Right)
			queue2 = append(queue2, top2.Left, top2.Right)
		} else if top1 == nil && top2 == nil {
			continue
		} else {
			return false
		}
	}
	return true
}