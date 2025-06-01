package binarySearch

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return res
	}
	first := find(nums, target, true)
	if first == -1 {
		return res
	}
	last := find(nums, target, false)
	return []int{first, last}

}

func find(nums []int, target int, first bool) int {
	left, right, result := 0, len(nums)-1, -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			result = mid
			if first {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return result
}
