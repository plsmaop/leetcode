/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max, index := -1, 0
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}
	newNode := TreeNode{
		Val: max,
	}
	newNode.Left = constructMaximumBinaryTree(nums[:index])
	newNode.Right = constructMaximumBinaryTree(nums[index+1:])
	return &newNode
}