func jump(nums []int) int {
    var (
        currEnd int         
        farthest int        
        jumps int      
    )
    
    for i:=0;i<len(nums) - 1;i++ {
        
        if (farthest < i + nums[i]) {              
            farthest = i + nums[i]  
        }
        if (i == currEnd) {                      
            jumps++                               
            currEnd = farthest               
        }
    }
    return jumps
}