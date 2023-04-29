func guessNumber(n int) int {
    low, high := 1, n
    for low <= high {
        mid1 :=  low + (high - low) / 3 
        mid2 := high - (high - low) / 3;
        
        res1 := guess(mid1)
        res2 := guess(mid2)
        if res1 == 0 {
            return mid1
        }
        if res2 == 0 {
            return mid2
        } else if res1 < 0 {
            high = mid1 - 1
        } else if res2 > 0 {
            low = mid2 + 1
        } else {
            low = mid1 + 1
            high = mid2 - 1
        }
    }
    return -1
}