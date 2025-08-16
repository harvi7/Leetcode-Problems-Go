package array

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	for _, interval := range intervals {
		if newInterval == nil || interval[1] < newInterval[0] {
			// Current interval ends before newInterval starts
			res = append(res, interval)
		} else if newInterval[1] < interval[0] {
			// newInterval ends before current interval starts
			res = append(res, newInterval)
			res = append(res, interval)
			newInterval = nil
		} else {
			// Overlapping intervals: merge them
			if interval[0] < newInterval[0] {
				newInterval[0] = interval[0]
			}
			if interval[1] > newInterval[1] {
				newInterval[1] = interval[1]
			}
		}
	}

	if newInterval != nil {
		res = append(res, newInterval)
	}

	return res
}
