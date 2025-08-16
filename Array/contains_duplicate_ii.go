package array

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int) // value → last index

	for i, current := range nums {
		if lastIndex, exists := indexMap[current]; exists && i-lastIndex <= k {
			return true
		}
		indexMap[current] = i
	}

	return false
}
