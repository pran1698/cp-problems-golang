package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	t := head
	l := 0
	for t != nil {
		l++
		t = t.Next
	}
	mid := l / 2
	var l2 *ListNode = nil
	t = head
	i := 0
	for t != nil {
		i++
		if i == mid {
			l2 = t.Next
			t.Next = nil
			break
		}
		t = t.Next
	}

	var prev *ListNode = nil
	for l2 != nil {
		t1 := l2.Next
		l2.Next = prev
		prev = l2
		l2 = t1
	}

	t = head
	i = 0
	for t != nil && prev != nil {
		t1 := t.Next
		t.Next = &ListNode{
			Val:  prev.Val,
			Next: t1,
		}
		prev = prev.Next
		t = t1
	}

	for prev != nil {
		t = head
		for t.Next != nil {
			t = t.Next
		}
		t.Next = prev
		prev = prev.Next
	}
}
