package twoPointers

func removeDuplicatesII(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	left := 2
	for right := 2; right < len(nums); right++ {
		if nums[right] != nums[left-2] {
			nums[left] = nums[right]
			left += 1
		}
	}
	return left
}
