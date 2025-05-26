func averageOfLevels(root *TreeNode) []float64 {
    result := []float64{}
    currLevel := []*TreeNode{root}
    for len(currLevel) > 0 {
        sum := 0
        nextLevel := []*TreeNode{}
        for _, node := range currLevel {
            sum = sum + node.Val
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel,node. Right)
            }
        }
        result = append(result, float64(sum)/float64(len(currLevel)))
        currLevel = nextLevel
    }
    return result 
}