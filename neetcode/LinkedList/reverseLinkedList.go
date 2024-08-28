package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var temp1 *ListNode = nil

	for head != nil {
		temp := head.Next
		head.Next = temp1
		temp1 = head
		head = temp
	}

	return temp1
}
