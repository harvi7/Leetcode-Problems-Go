func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    prehead := &ListNode{}

    prev := prehead
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            prev.Next = l1
            l1 = l1.Next
        } else {
            prev.Next = l2
            l2 = l2.Next
        }
        prev = prev.Next
    }

    if l1 == nil {
        prev.Next = l2
    } else {
        prev.Next = l1
    }

    return prehead.Next;
}