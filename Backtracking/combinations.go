func combine(n int, k int) [][]int {
    result := [][]int{}
    if k > n {
        return result
    }
    helper([]int{}, n, k, 1, &result)
    return result
}

func helper(currentList []int, n, k, start int, result *[][]int) {
    if k == 0 {
        tmp := make([]int, len(currentList))
        copy(tmp, currentList)
        *result = append(*result, tmp)
    }
    
    for i := start; i <= n; i++ {
        currentList = append(currentList, i) 
        helper(currentList, n, k - 1, i + 1, result)
        currentList = currentList[:len(currentList) - 1]
    }
}