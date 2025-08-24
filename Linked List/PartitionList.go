package linkedlist

// https://www.youtube.com/watch?v=rKnD7rLT0lI&ab_channel=NeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{0, nil}
	before := beforeHead
	afterHead := &ListNode{0, nil}
	after := afterHead
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}
