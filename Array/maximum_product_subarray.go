package array

import "../helper"

func maxProduct(nums []int) int {
	curMax, curMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		temp := curMax
		curMax = helper.Max(nums[i], helper.Max(temp*nums[i], curMin*nums[i]))
		curMin = helper.Min(nums[i], helper.Min(temp*nums[i], curMin*nums[i]))
		if curMax > ans {
			ans = curMax
		}
	}
	return ans
}
