/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	type node struct {
		node  *TreeNode
		level int
	}
	var tmp node
	tmp.node = root
	tmp.level = 1
	queue := []node{tmp}

	for len(queue) > 0 {
		top := queue[0]
		if top.node.Left == nil && top.node.Right == nil {
			ans = top.level
			break
		}
		if top.node.Left != nil {
			var ltmp node
			ltmp.node = top.node.Left
			ltmp.level = top.level + 1
			queue = append(queue, ltmp)
		}
		if top.node.Right != nil {
			var rtmp node
			rtmp.node = top.node.Right
			rtmp.level = top.level + 1
			queue = append(queue, rtmp)
		}
		queue = queue[1:]
	}
	return ans
}