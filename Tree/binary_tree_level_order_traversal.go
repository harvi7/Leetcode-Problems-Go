func levelOrder(root *TreeNode) [][]int {
    var result [][]int
    
    if root == nil { return result }
    
    var queue []*TreeNode
    enqueue(&queue, root)
    
    for len(queue) > 0 {
        queueSize := len(queue)
        var levelNodes []int
        
        for i := 0; i < queueSize; i++ {
            dequeuedEle := dequeue(&queue)
            
            levelNodes = append(levelNodes, dequeuedEle.Val)
            
            if dequeuedEle.Left != nil {
                enqueue(&queue, dequeuedEle.Left)
            }
            
            if dequeuedEle.Right != nil {
                enqueue(&queue, dequeuedEle.Right)
            }
        }
        
        result = append(result, levelNodes)
    }

    return result
}

func enqueue(queue *[]*TreeNode, item *TreeNode) {
    if queue == nil { panic("nil pointer") }
    
    *queue = append(*queue, item)
} 

func dequeue(queue *[]*TreeNode) *TreeNode {
    if queue == nil { panic("nil pointer") }
    
    dequeuedElement := (*queue)[0]
    
    *queue = (*queue)[1:]
    
    return dequeuedElement
}