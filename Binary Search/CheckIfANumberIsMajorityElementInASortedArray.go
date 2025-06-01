package binarySearch

// https://leetcode.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array/editorial/

func isMajorityElement(nums []int, target int) bool {
	firstOccurrenceIndex := firstOccurrence(nums, target)
	if firstOccurrenceIndex == -1 {
		return false
	}
	lastOccurrenceIndex := lastOccurrence(nums, target)
	if lastOccurrenceIndex-firstOccurrenceIndex+1 > len(nums)/2 {
		return true
	}
	return false
}

func firstOccurrence(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if (mid == 0 || target > nums[mid-1]) && nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func lastOccurrence(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if (mid == len(nums)-1 || target < nums[mid+1]) && nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
