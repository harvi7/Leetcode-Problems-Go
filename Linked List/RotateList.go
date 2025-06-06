func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil{
        return head
    }
    var dummy  = &ListNode{Next: head}
    count, tail, sp := 0, head, head
    
    for tail.Next != nil {
        count++
        tail = tail.Next
    }
    count++
    j := k % count
    for i := 1; i < count - j; i++ {
        sp = sp.Next
    }
    
    tail.Next = head
    dummy.Next = sp.Next
    sp.Next = nil
    
    return dummy.Next
}