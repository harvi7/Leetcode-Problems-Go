package array

import (
	"math"

	"../helper"
)

func maxSubarraySumCircular(nums []int) int {
	curMin, minSum := math.MaxInt32, math.MaxInt32
	curMax, maxSum := -math.MaxInt32, -math.MaxInt32
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		curMin = helper.Min(nums[i], curMin+nums[i])
		curMax = helper.Max(nums[i], curMax+nums[i])
		minSum = helper.Min(minSum, curMin)
		maxSum = helper.Max(maxSum, curMax)
	}
	if total == minSum {
		return maxSum
	}
	return max(total-minSum, maxSum)
}
