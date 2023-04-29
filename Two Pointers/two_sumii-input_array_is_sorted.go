func twoSum(numbers []int, target int) []int {
    low, high := 0, len(numbers) - 1
    for low < high {
        s := numbers[low] + numbers[high]
        if s == target {
            return []int{low + 1, high + 1}
        } else if s < target {
            low += 1
        } else {
            high -= 1
        }
    }
    return []int{-1, -1}
}