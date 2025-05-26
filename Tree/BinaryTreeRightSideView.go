func rightSideView(root *TreeNode) []int {
    levels := []int{}

    var walk func(n *TreeNode, depth int)
    walk = func(n *TreeNode, d int) {
        if n == nil {
            return
        }
        if len(levels) < d + 1 {
            levels = append(levels, n.Val)
        }
        walk(n.Right, d + 1)
        walk(n.Left, d + 1)
        return
    }

    walk(root, 0)
    return levels
}