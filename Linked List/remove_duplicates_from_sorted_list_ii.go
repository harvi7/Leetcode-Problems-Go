package linkedlist

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // dummy helps handle head duplicates cleanly
	prev := dummy                  // prev points to last node in result list
	curr := head

	for curr != nil {
		// Detect duplicates
		duplicate := false
		for curr.Next != nil && curr.Val == curr.Next.Val {
			duplicate = true
			curr = curr.Next
		}

		if duplicate {
			// Skip all duplicates
			prev.Next = curr.Next
		} else {
			// No duplicate, move prev
			prev = prev.Next
		}

		curr = curr.Next
	}

	return dummy.Next
}
