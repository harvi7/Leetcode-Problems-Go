package array

// https://leetcode.com/problems/intersection-of-three-sorted-arrays/editorial/

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	var ans []int
	var p1, p2, p3 int
	len1, len2, len3 := len(arr1), len(arr2), len(arr3)
	for p1 < len1 && p2 < len2 && p3 < len3 {
		if arr1[p1] == arr2[p2] && arr2[p2] == arr3[p3] {
			ans = append(ans, arr1[p1])
			p1++
			p2++
			p3++
		} else {
			if arr1[p1] < arr2[p2] {
				p1++
			} else if arr2[p2] < arr3[p3] {
				p2++
			} else {
				p3++
			}
		}
	}
	return ans
}
