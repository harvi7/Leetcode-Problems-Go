package twoPointers

func removeDuplicates(nums []int) int {
	prev := nums[0]
	insertIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
		prev = nums[i]
	}
	return insertIndex
}
