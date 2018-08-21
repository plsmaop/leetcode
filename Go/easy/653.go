/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func merge(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	left := merge(node.Left)
	right := merge(node.Right)
	array := append(left, node.Val)
	array = append(array, right...)
	return array
}

func findTarget(root *TreeNode, k int) bool {
	array := merge(root)
	start, end := 0, len(array)-1
	for start+end < 2*len(array) && start != end {
		tmp := array[start] + array[end]
		if tmp == k {
			return true
		}
		if tmp > k {
			end--
		} else {
			start++
		}
	}
	return false
}