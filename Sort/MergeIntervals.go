package sort

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sortIntervals(intervals)

	finalSlice := [][]int{}

	start, end := intervals[0][0], intervals[0][1]

	for idx := 1; idx < len(intervals); idx++ {
		currentStart := intervals[idx][0]
		currentEnd := intervals[idx][1]

		if currentStart <= end {
			if currentEnd > end {
				end = currentEnd
			}
		} else {
			finalSlice = append(finalSlice, []int{start, end})
			start = currentStart
			end = currentEnd
		}
	}

	finalSlice = append(finalSlice, []int{start, end})

	return finalSlice
}

func sortIntervals(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}
