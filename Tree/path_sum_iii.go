func pathSum(root *TreeNode, targetSum int) int {
    h := make(map[int]int)
    count := 0
    h[0] = 1
    
    preorder(root, h, 0, targetSum, &count)
    return count
}

func preorder(root *TreeNode, h map[int]int, currSum int, sum int, count *int) {
    if root == nil {
        return
    }
    currSum += root.Val

    *count += h[currSum - sum]

    h[currSum]++
    preorder(root.Left, h, currSum, sum, count)
    preorder(root.Right, h, currSum, sum, count)
    h[currSum]--
}