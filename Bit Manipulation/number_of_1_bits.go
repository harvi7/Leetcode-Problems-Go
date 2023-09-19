func hammingWeight(num uint32) int {
    res := 0;
    for i := 0; i < 32; i++ {
        if num & uint32(1) == 1 {
            res += 1    
        }
        num = num >> 1
    }
    return res;
}