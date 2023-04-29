func nthUglyNumber(n int) int {
    nums := make([]int, n)
	nums[0] = 1
    index2, index3, index5 := 0, 0, 0
    min_num := 0
    for i := 1; i < n; i++ {
        min_num = Min(nums[index2] * 2, Min(nums[index3] * 3, nums[index5] * 5))
        if min_num == nums[index2] * 2 { index2 += 1 }
        if min_num == nums[index3] * 3 { index3 += 1 }
        if min_num == nums[index5] * 5 { index5 += 1 }
        nums[i] = min_num
    }
    return nums[n - 1]
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}