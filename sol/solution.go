package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	result := InOrderTraversalKth(root, k, &count)
	return result.Val
}

func InOrderTraversalKth(root *TreeNode, k int, count *int) *TreeNode {
	if root == nil { // last node
		return nil
	}
	left := InOrderTraversalKth(root.Left, k, count)
	if *count == k {
		return left
	}
	*count++
	if *count == k {
		return root
	}
	return InOrderTraversalKth(root.Right, k, count)
}
