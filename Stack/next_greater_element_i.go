func nextGreaterElement(nums1 []int, nums2 []int) []int {
    var stack []int
    nextGreatest := make(map[int]int)
    for i := 0; i < len(nums2); i++ {
        for len(stack) != 0 {
            a := stack[len(stack) - 1]
            if nums2[i] <= a {
                break
            } else {
                nextGreatest[a] = nums2[i]
                stack = stack[:len(stack) - 1]
            }
        }
        stack = append(stack, nums2[i])
    }

    res := make([]int, len(nums1))
    for i := 0; i < len(nums1); i++ {
        if v, ok := nextGreatest[nums1[i]]; ok {
            res[i] = v
        } else {
            res[i] = -1
        }
    }
    return res
}