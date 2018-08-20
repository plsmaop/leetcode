import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(node *TreeNode, ans *string) {
  if node == nil {
    return
  }
  *ans += fmt.Sprintf("%d", node.Val)
  if node.Left == nil && node.Right == nil {
    return
  }
  *ans += "("
  dfs(node.Left, ans)
  *ans += ")"
  if node.Right == nil {
    return
  }
  *ans += "("
  dfs(node.Right, ans)
  *ans += ")"
}

func tree2str(t *TreeNode) string {
  if t == nil {
    return ""
  }
  ans := ""
  dfs(t, &ans)
  return ans
}