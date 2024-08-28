package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	t := head

	i := 0
	var prev *ListNode = nil
	for t != nil && t.Next != nil {
		t1 := t.Next
		t.Next = t1.Next
		t1.Next = t
		if i == 0 {
			head = t1
		} else {
			prev.Next = t1
		}
		prev = t
		i++

		t = t.Next
	}

	return head
}
