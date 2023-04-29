func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n <= 0 {
        return head
    }
        
    dummy := &ListNode{Next: head}
    dummy.Next = head
    preDelete := dummy

    for i := 0; i < n; i++ {
        if head == nil {
            return nil
        }
        head = head.Next
    }

    for head != nil {
        preDelete = preDelete.Next
        head = head.Next
    }
    preDelete.Next = preDelete.Next.Next

    return dummy.Next
}