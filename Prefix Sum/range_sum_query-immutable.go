type NumMatrix struct {
    preSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    preSumMatrix := NumMatrix{}
    m, n := len(matrix), len(matrix[0])
    preSumMatrix.preSum = make([][]int, m)
    for i := 0; i < m; i++ {
        preSumMatrix.preSum[i] = make([]int, n)
    }
    preSumMatrix.preSum[0][0] = matrix[0][0]
    for j := 1; j < n; j++ {
        preSumMatrix.preSum[0][j] = preSumMatrix.preSum[0][j - 1] + matrix[0][j]
    }
    for i := 1; i < m; i++ {
        preSumMatrix.preSum[i][0] = preSumMatrix.preSum[i - 1][0] + matrix[i][0]
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            preSumMatrix.preSum[i][j] = preSumMatrix.preSum[i - 1][j] + preSumMatrix.preSum[i][j - 1] - preSumMatrix.preSum[i - 1][j - 1] + matrix[i][j]
        }
    }
    return preSumMatrix
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    a1 := this.preSum[row2][col2]
    var a2, a3, a4 int
    if row1 > 0 {
        a2 = this.preSum[row1 - 1][col2]
    }
    if col1 > 0 {
        a3 = this.preSum[row2][col1 - 1]
    }
    if row1 > 0 && col1 > 0 {
        a4 = this.preSum[row1 - 1][col1 - 1]
    }
    return a1 - a2 - a3 + a4
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */