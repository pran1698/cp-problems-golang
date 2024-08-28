package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	s := head
	f := head.Next

	for s != nil && f != nil {
		if s == f {
			return true
		}

		s = s.Next
		i := 0
		for f != nil && i < 2 {
			f = f.Next
			i++
		}
	}

	return false
}
