package test

import "testing"

type ListNode25 struct {
	Val  int
	Next *ListNode25
}

func reverseKGroup(head *ListNode25, k int) *ListNode25 {
	if head == nil || k <= 1 {
		return head
	}
	var (
		result *ListNode25 = nil
		gPre   *ListNode25 = nil
	)
	ptr, i, gFirst := head, 0, head
	for ptr != nil {
		ptrNext := ptr.Next
		if i == k-1 { // 到达一组的最后一个
			gh, gt := reverse(gFirst, k)
			if result == nil {
				result = gh
			}
			if gPre != nil {
				gPre.Next = gh
			}
			gPre = gt
			gFirst = ptrNext
			i = 0
		} else {
			i += 1
		}
		ptr = ptrNext
	}
	return result
}

func reverse(node *ListNode25, k int) (head *ListNode25, tail *ListNode25) {
	tail, head = node, node
	var (
		pre  *ListNode25 = nil
		next *ListNode25 = nil
	)
	for i := 0; i < k; i++ {
		next = head.Next
		if pre != nil {
			head.Next = pre
		}
		pre = head
		if i < k-1 {
			head = next
		}
	}
	tail.Next = next
	return head, tail
}

func TestDo25(t *testing.T) {
	head := &ListNode25{
		Val: 1,
		Next: &ListNode25{
			Val: 2,
			Next: &ListNode25{
				Val: 3,
				Next: &ListNode25{
					Val: 4,
					Next: &ListNode25{
						Val: 5,
						Next: &ListNode25{
							Val:  6,
							Next: &ListNode25{Val: 7}},
					},
				}}}}
	result := reverseKGroup(head, 3)
	println(result)
}
