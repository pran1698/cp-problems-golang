package main

func rightSideView(root *TreeNode) []int {
	var q []*TreeNode
	var res []int

	enqueue(&q, root)
	for len(q) > 0 {
		l := len(q)
		var levelNodes []int
		for i := 0; i < l; i++ {
			n := dequeue(&q)
			levelNodes = append(levelNodes, n.Val)
			enqueue(&q, n.Left)
			enqueue(&q, n.Right)
		}
		res = append(res, levelNodes[len(levelNodes)-1])
	}

	return res
}
