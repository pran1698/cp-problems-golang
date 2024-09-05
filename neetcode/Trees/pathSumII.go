package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traversalSum(root *TreeNode, targetSum, sum int, curr []int, ans *[][]int) {
	if root == nil {
		return
	}

	sum += root.Val
	curr = append(curr, root.Val)
	if targetSum == sum && root.Left == nil && root.Right == nil {
		path := make([]int, len(curr))
		copy(path, curr)
		*ans = append(*ans, path)
	}

	traversalSum(root.Left, targetSum, sum, curr, ans)
	traversalSum(root.Right, targetSum, sum, curr, ans)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	traversalSum(root, targetSum, 0, []int{}, &ans)
	return ans
}
