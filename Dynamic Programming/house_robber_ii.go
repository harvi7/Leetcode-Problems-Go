package dp

import "math"

func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return int(math.Max(float64(rob_simple(nums, 0, len(nums)-2)), float64(rob_simple(nums, 1, len(nums)-1))))
}

func rob_simple(nums []int, l, r int) int {
	var ret, pre int
	for i := l; i <= r; i++ {
		ret, pre = int(math.Max(float64(ret), float64(pre+nums[i]))), ret
	}
	return ret
}
