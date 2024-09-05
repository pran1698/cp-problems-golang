package main

func heightTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(heightTree(root.Left), heightTree(root.Right))
}

func levelOrderTraversal(root *TreeNode, index, level int, ans [][]int) {
	if root == nil {
		return
	}

	if level == 1 {
		ans[index-1] = append(ans[index-1], root.Val)
	} else if level > 1 {
		levelOrderTraversal(root.Left, index, level-1, ans)
		levelOrderTraversal(root.Right, index, level-1, ans)
	}
}

func levelOrder(root *TreeNode) [][]int {
	height := heightTree(root)

	var ans = make([][]int, height)
	for i := 1; i <= height; i++ {
		levelOrderTraversal(root, i, i, ans)
	}

	return ans
}
