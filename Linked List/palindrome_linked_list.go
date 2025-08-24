package linkedlist

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		nxt := slow.Next
		slow.Next = prev
		prev = slow
		slow = nxt
	}

	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}
	return true
}
