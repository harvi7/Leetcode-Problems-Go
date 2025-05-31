package array

// https://leetcode.com/problems/rotate-image/

func rotate(matrix [][]int) {
	n := len(matrix)
	transpose(matrix, n)
	reflect(matrix, n)
}

func transpose(matrix [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}

func reflect(matrix [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][n-j-1]
			matrix[i][n-j-1] = tmp
		}
	}
}
