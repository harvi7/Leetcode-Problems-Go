func containsDuplicate(nums []int) bool {
	set := make(map[int]int, len(nums))
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		}
		set[n] = 1
	}
	return false
}