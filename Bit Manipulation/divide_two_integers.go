package bitManipulation

import "math"

// https://www.youtube.com/watch?v=htX69j1jf5U&ab_channel=NideeshTerapalli

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	isNegative := (dividend < 0) != (divisor < 0)

	a := int64(math.Abs(float64(dividend)))
	b := int64(math.Abs(float64(divisor)))

	var res int64 = 0

	for a-b >= 0 {
		var x int64 = 0
		for a-(b<<(x+1)) >= 0 {
			x++
		}
		res += 1 << x
		a -= b << x
	}
	if isNegative {
		return int(-res)
	}
	return int(res)
}
