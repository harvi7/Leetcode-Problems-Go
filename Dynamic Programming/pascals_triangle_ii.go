func getRow(rowIndex int) []int {
    result := []int{}

    for i := 0; i <= rowIndex; i++ {
        temp := make([]int, len(result))
        copy(temp, result)
        result = append(result, 1)
        for j := 1; j < len(result) - 1; j++ {
            result[j] = temp[j - 1] + temp[j]
        }
    }
    return result
}