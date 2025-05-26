package array

func minimumSwaps(nums []int) int {
	smallest, smallestIdx := nums[0], 0
	largest, largestIdx := nums[0], 0
	for i, val := range nums {
		// Find the smallest element with smallest idx
		if smallest > val {
			smallest = val
			smallestIdx = i
		}
		// Find the greatest element with greatest idx
		if largest <= val {
			largest = val
			largestIdx = i
		}
	}

	res := smallestIdx + (len(nums) - largestIdx - 1)
	if smallestIdx > largestIdx {
		res--
	}
	return res
}
