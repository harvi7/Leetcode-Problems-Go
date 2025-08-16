package twopointers

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum, result, left := 0, math.MaxInt32, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			result = min(result, i+1-left)
			sum -= nums[left]
			left++
		}
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result
}
