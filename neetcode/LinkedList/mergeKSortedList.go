package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{
		Val:  0,
		Next: nil,
	}
	t := h

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			t.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			t = t.Next
			l1 = l1.Next
		} else {
			t.Next = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}
			t = t.Next
			l2 = l2.Next
		}
	}

	for l1 != nil {
		t.Next = &ListNode{
			Val:  l1.Val,
			Next: nil,
		}
		t = t.Next
		l1 = l1.Next
	}

	for l2 != nil {
		t.Next = &ListNode{
			Val:  l2.Val,
			Next: nil,
		}
		t = t.Next
		l2 = l2.Next
	}

	return h.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	for len(lists) > 1 {
		list1 := lists[0]
		list2 := lists[1]

		lt := mergeTwoSortedList(list1, list2)
		lists = lists[2:]
		lists = append(lists, lt)
	}

	if len(lists) > 0 {
		return lists[0]
	}
	return nil
}
