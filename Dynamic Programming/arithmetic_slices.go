package dp

func numberOfArithmeticSlices(nums []int) int {
	dp, result := 0, 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp = 1 + dp
			result += dp
		} else {
			dp = 0
		}
	}
	return result
}
