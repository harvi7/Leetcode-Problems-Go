package greedy

func canJump(nums []int) bool {
	lastGoodIndexPosition := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastGoodIndexPosition {
			lastGoodIndexPosition = i
		}
	}

	return lastGoodIndexPosition == 0
}
