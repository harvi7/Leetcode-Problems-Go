package array

func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	longestStreak := 0

	for num := range numSet {
		// Check if it's the start of a sequence
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			currentStreak := 1

			for {
				if _, exists := numSet[currentNum+1]; exists {
					currentNum++
					currentStreak++
				} else {
					break
				}
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
