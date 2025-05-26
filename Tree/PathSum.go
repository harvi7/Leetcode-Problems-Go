func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    return helper(root, targetSum)
}
    
func helper(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    sum -= root.Val
    if root.Left == nil && root.Right == nil && sum == 0 {
        return true
    }
    return helper(root.Left, sum) || helper(root.Right, sum)
}