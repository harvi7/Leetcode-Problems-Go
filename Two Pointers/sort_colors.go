package twopointers

func sortColors(nums []int) {
	one := 0
	two := len(nums) - 1
	for i := 0; i <= two; {
		if nums[i] == 2 {
			nums[two], nums[i] = nums[i], nums[two]
			two--
		} else if nums[i] == 0 {
			nums[one], nums[i] = nums[i], nums[one]
			one++
			i++
		} else {
			i++
		}
	}
}
