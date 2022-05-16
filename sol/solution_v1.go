package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallestV1(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	cur := root
	for {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		topIdx := len(stack) - 1
		cur = stack[topIdx]
		stack = stack[:topIdx]
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
}
