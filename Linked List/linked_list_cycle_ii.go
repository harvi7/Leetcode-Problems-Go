package linkedlist

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head

	// Step 1: Detect cycle using Floyd’s Tortoise and Hare
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// Step 2: Find the entry point of cycle
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
