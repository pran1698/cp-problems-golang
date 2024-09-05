package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getIndex(inorder []int, v int) int {
	for i, n := range inorder {
		if n == v {
			return i
		}
	}

	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var ans *TreeNode = &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	i := getIndex(inorder, preorder[0])
	//fmt.Println(i, preorder[0])
	ans.Left = buildTree(preorder[1:i+1], inorder[0:i])
	ans.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return ans
}
