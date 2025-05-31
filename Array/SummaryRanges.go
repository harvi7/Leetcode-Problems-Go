package array

import "fmt"

func summaryRanges(nums []int) []string {
	result := []string{}

	if len(nums) == 0 {
		return []string{}
	}
	left := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-prev != 1 {
			result = append(result, fmtString(left, prev))
			left, prev = nums[i], nums[i]
		} else {
			prev = nums[i]
		}
	}
	return append(result, fmtString(left, prev))
}

func fmtString(l, r int) string {
	if l == r {
		return fmt.Sprint(l)
	}

	return fmt.Sprint(l, "->", r)
}
