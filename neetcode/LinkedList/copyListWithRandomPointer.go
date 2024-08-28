package main

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func copyRandomList(head *RandomListNode) *RandomListNode {
	m := make(map[*RandomListNode]*RandomListNode)
	t := head

	for t != nil {
		n := &RandomListNode{
			Val:    t.Val,
			Next:   nil,
			Random: nil,
		}
		m[t] = n
		t = t.Next
	}

	res := &RandomListNode{
		Val:    0,
		Next:   nil,
		Random: nil,
	}
	t1 := res
	t = head
	for t != nil {
		curr := m[t]
		next := m[t.Next]
		random := m[t.Random]

		curr.Next = next
		curr.Random = random
		t1.Next = curr
		t1 = t1.Next
		t = t.Next
	}

	return res.Next
}
