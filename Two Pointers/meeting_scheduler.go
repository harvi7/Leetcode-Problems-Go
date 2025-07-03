package twopointers

import "sort"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})

	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})

	pointer1, pointer2 := 0, 0

	for pointer1 < len(slots1) && pointer2 < len(slots2) {
		var intersectLeft = max(slots1[pointer1][0], slots2[pointer2][0])
		var intersectRight = -1 * max(-slots1[pointer1][1], -slots2[pointer2][1])

		if (intersectRight - intersectLeft) >= duration {
			return []int{intersectLeft, intersectLeft + duration}
		}

		// Increment the point as needed
		if slots1[pointer1][1] < slots2[pointer2][1] {
			pointer1 += 1
		} else {
			pointer2 += 1
		}
	}
	return []int{}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
