package main

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resHead := &ListNode{Val: 0, Next: nil}
	tempHead := resHead
	var ans, carry int

	for l1 != nil && l2 != nil {
		ans = l1.Val + l2.Val + carry

		node := &ListNode{}
		node.Val = ans % 10
		tempHead.Next = node
		carry = ans / 10
		tempHead = tempHead.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		ans = l1.Val + carry

		node := &ListNode{}
		node.Val = ans % 10
		tempHead.Next = node
		carry = ans / 10

		tempHead = tempHead.Next
		l1 = l1.Next
	}

	for l2 != nil {
		ans = l2.Val + carry

		node := &ListNode{}
		node.Val = ans % 10
		tempHead.Next = node
		carry = ans / 10

		tempHead = tempHead.Next
		l2 = l2.Next
	}

	if carry != 0 {
		node := &ListNode{}
		node.Val = carry
		tempHead.Next = node
		tempHead = tempHead.Next
	}

	return resHead.Next
}
