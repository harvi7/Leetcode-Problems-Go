func bstFromPreorder(preorder []int) *TreeNode {
    if len(preorder) == 0 {return nil}
    
    root := &TreeNode{preorder[0], nil, nil}
    s := []*TreeNode{root}
    for i := 1; i < len(preorder); i++ {
        node := s[len(s) - 1]
        child := &TreeNode{preorder[i], nil, nil}
        
        for len(s) > 0 && s[len(s) - 1].Val < child.Val {
            node = s[len(s) - 1]
            s = s[:len(s) - 1]
        }
        
        if node.Val < child.Val {
            node.Right = child
        }else{
            node.Left = child
        }
        s = append(s, child)
    }
    return root
}