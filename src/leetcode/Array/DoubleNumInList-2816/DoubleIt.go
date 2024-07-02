package DoubleNumInList_2816

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	h := rev(head)
	c := 0
	var pre *ListNode = h
	for p := h; p != nil; p = p.Next {
		p.Val = p.Val*2 + c
		c = p.Val / 10
		p.Val = p.Val % 10
		pre = p
	}
	if c > 0 {
		t := &ListNode{Val: c}
		pre.Next = t
	}
	return rev(h)
}

func rev(h *ListNode) (newHead *ListNode) {
	if h == nil || h.Next == nil {
		return h
	}
	newHead = rev(h.Next)
	h.Next.Next = h
	h.Next = nil
	return newHead
}
