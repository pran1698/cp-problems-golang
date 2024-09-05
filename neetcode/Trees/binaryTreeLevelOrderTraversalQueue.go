package main

func enqueue(q *[]*TreeNode, n *TreeNode) {
	if n == nil {
		return
	}
	*q = append(*q, n)
}

func dequeue(q *[]*TreeNode) *TreeNode {
	if q == nil {
		return nil
	}
	dequeuedElement := (*q)[0]
	*q = (*q)[1:]
	return dequeuedElement
}

func levelOrderQueue(root *TreeNode) [][]int {
	var res [][]int
	var q []*TreeNode

	enqueue(&q, root)

	for len(q) > 0 {
		var levelNodes []int
		l := len(q)
		for i := 0; i < l; i++ {
			n := dequeue(&q)
			levelNodes = append(levelNodes, n.Val)
			enqueue(&q, n.Left)
			enqueue(&q, n.Right)
		}
		res = append(res, levelNodes)
	}

	return res
}
