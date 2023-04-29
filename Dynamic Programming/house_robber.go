func rob(nums []int) int {    
    prevMax, currMax := 0, 0
    
    for i:=0; i < len(nums); i++ {
        temp := currMax
        if prevMax + nums[i] > currMax {
            currMax = prevMax + nums[i] 
        }
        prevMax = temp
    }
    return currMax
}