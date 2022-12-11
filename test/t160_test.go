package test

import (
	"testing"
)

type ListNode160 struct {
	Val  int
	Next *ListNode160
}

func getIntersectionNode(headA, headB *ListNode160) *ListNode160 {
	cache, p := make(map[int]*ListNode160), headA
	for p != nil {
		cache[p.Val] = p
		p = p.Next
	}
	p = headB
	for p != nil {
		if node, ok := cache[p.Val]; ok {
			return node
		}
	}
	return nil
}

func TestDo160(t *testing.T) {

}
