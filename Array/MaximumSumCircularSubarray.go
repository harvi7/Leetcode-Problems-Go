package array

import "math"

func maxSubarraySumCircular(nums []int) int {
	curMin, minSum := math.MaxInt32, math.MaxInt32
	curMax, maxSum := -math.MaxInt32, -math.MaxInt32
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		curMin = min(nums[i], curMin+nums[i])
		curMax = max(nums[i], curMax+nums[i])
		minSum = min(minSum, curMin)
		maxSum = max(maxSum, curMax)
	}
	if total == minSum {
		return maxSum
	}
	return max(total-minSum, maxSum)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
