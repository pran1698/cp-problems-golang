package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	t := head
	t1 := head
	i := 0
	for i < n {
		i++
		t1 = t1.Next
	}
	if t1 == nil {
		return head.Next
	}

	for t1.Next != nil {
		t = t.Next
		t1 = t1.Next
	}

	t.Next = t.Next.Next
	return head
}
