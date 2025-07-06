package hashtable

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	for j := 0; j < len(nums2); j++ {
		if val, ok := m[nums2[j]]; ok {
			if val >= 1 {
				m[nums2[j]] -= 1
				res = append(res, nums2[j])
			}
		}
	}
	return res
}
