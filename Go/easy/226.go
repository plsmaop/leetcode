/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	array := []*TreeNode{root}
	for len(array) > 0 {
		if array[0].Left != nil {
			array = append(array, array[0].Left)
		}
		if array[0].Right != nil {
			array = append(array, array[0].Right)
		}
		array[0].Left, array[0].Right = array[0].Right, array[0].Left
		array = array[1:]
	}
	return root
}