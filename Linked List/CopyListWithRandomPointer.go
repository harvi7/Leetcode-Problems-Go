package linkedlist

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	curr := head
	for curr != nil {
		curr, curr.Next = curr.Next, &RandomNode{Val: curr.Val, Next: curr.Next}
	}

	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	var dummyHead *RandomNode
	if head != nil {
		dummyHead = head.Next
	}
	curr = head
	for curr != nil {
		copyNode := curr.Next
		oldNext := copyNode.Next
		copyNode.Next = nil
		if oldNext != nil {
			copyNode.Next = oldNext.Next
		}
		curr.Next = oldNext
		curr = curr.Next
	}
	return dummyHead
}
