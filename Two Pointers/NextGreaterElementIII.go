package twopointers

import (
	"math"
	"sort"
)

func nextGreaterElement(n int) int {
	nums := numToSlice(n)
	if len(nums) == 1 {
		return -1
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] { // we found an "inversion", we only need to find one
			for j := len(nums) - 1; j >= i; j-- { // look for the first digit higher than the one found in the inversion
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					temp := nums[i:]
					sort.Ints(temp)
					nums = append(nums[:i], temp...)
					res := sliceToNum(nums)
					if res > math.MaxInt32 {
						return -1
					}
					return res
				}
			}
		}
	}
	return -1
}

func numToSlice(in int) []int {
	var nums []int
	for in > 0 {
		digit := in % 10
		nums = append([]int{digit}, nums...)
		in = in / 10
	}
	return nums
}

func sliceToNum(in []int) int {
	base := 1
	result := 0
	for i := len(in) - 1; i >= 0; i-- {
		result += base * in[i]
		base = base * 10
	}
	return result
}
