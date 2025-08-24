package array

func setZeroes(matrix [][]int) {
	isColZero := false
	R := len(matrix)
	C := len(matrix[0])

	// Step 1: use first row/col as markers
	for i := 0; i < R; i++ {
		if matrix[i][0] == 0 {
			isColZero = true
		}
		for j := 1; j < C; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// Step 2: zero out based on markers
	for i := 1; i < R; i++ {
		for j := 1; j < C; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Step 3: handle first row
	if matrix[0][0] == 0 {
		for j := 0; j < C; j++ {
			matrix[0][j] = 0
		}
	}

	// Step 4: handle first column
	if isColZero {
		for i := 0; i < R; i++ {
			matrix[i][0] = 0
		}
	}
}
