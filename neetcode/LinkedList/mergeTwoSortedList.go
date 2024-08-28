package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: list1,
	}

	var prev *ListNode = res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev = list1
			list1 = list1.Next
		} else {
			prev.Next = &ListNode{
				Val:  list2.Val,
				Next: list1,
			}
			prev = prev.Next
			list2 = list2.Next
		}
	}

	for list2 != nil {
		prev.Next = &ListNode{
			Val:  list2.Val,
			Next: list1,
		}
		prev = prev.Next
		list2 = list2.Next
	}

	return res.Next
}
