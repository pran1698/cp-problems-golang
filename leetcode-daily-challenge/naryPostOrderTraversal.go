package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func rec(root *Node, ans []int) []int {
	if root == nil {
		return ans
	}
	for _, n := range root.Children {
		ans = rec(n, ans)
	}

	ans = append(ans, root.Val)
	return ans
}

func postorder(root *Node) []int {
	ans := make([]int, 0)
	return rec(root, ans)
}
