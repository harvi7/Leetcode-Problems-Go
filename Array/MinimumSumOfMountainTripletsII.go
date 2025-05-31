package array

import "math"

// https://www.youtube.com/watch?v=GegaVi1CeUM

func minimumSum(nums []int) int {
	minR := make([]int, len(nums))
	minR[len(minR)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 2; i-- {
		minR[i] = min(nums[i], minR[i+1])
	}

	result := math.MaxInt32
	minL := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > minL && nums[i] > minR[i+1] {
			result = min(result, nums[i]+minL+minR[i+1])
		}
		minL = min(minL, nums[i])
	}

	if result == math.MaxInt32 {
		return -1
	}

	return result
}
