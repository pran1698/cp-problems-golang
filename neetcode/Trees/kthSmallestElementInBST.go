package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traversal(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, ans)
	*ans = append(*ans, root.Val)
	traversal(root.Right, ans)
}

func kthSmallest(root *TreeNode, k int) int {
	var ans []int
	traversal(root, &ans)
	return ans[k-1]
}
