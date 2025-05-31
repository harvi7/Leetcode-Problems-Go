package array

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/editorial/

func countNegatives(grid [][]int) int {
	n := len(grid[0])
	count, currRowNegativeIndex := 0, n-1
	for _, row := range grid {
		for currRowNegativeIndex >= 0 && row[currRowNegativeIndex] < 0 {
			currRowNegativeIndex--
		}
		count += (n - (currRowNegativeIndex + 1))
	}
	return count
}
