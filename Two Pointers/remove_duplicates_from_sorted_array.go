func removeDuplicates(nums []int) int {
    prev := nums[0]
    insert_index := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != prev {
            nums[insert_index] = nums[i]
            insert_index++
        } 
        prev = nums[i]
    }
    return insert_index
}