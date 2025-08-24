package linkedlist

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
