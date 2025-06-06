package twoPointers

func moveZeroes(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}

	left, right := 0, 0

	for right < n {
		if nums[right] == 0 {
			right++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left += 1
			right += 1
		}
	}
}
