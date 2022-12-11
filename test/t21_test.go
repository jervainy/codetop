package test

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := list1
	for {
		if list2 == nil {
			break
		}
		if list1.Next == nil {
			break
		} else {
			if list1.Val < list2.Val {
				list1 = list1.Next
			} else {
				//next := list2.Next
				list2.Next = list1
			}
			list1 = list1.Next
		}
	}
	return head
}
