package math

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	l, r, sum := 0, len(nums)-1, 0
	for l < r {
		sum += nums[r] - nums[l]
		l++
		r--
	}
	return sum
}
