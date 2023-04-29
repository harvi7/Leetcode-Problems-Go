func judgeSquareSum(c int) bool {
    for a := 0; a * a <= c; a++ {
        if isSquare(c - (a * a)) {
            return true
        }
    }
    return false
}

func isSquare(n int) bool {
    squareRooted := int(math.Sqrt(float64(n)))
    
    if squareRooted * squareRooted == n {
        return true
    }
    return false
}