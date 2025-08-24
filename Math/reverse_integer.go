package math

import "math"

func reverse(x int) int {
	signMultiplier := 1
	if x < 0 {
		signMultiplier = -1
		x = signMultiplier * x
	}

	var rev int
	for x > 0 {
		rev = rev*10 + x%10
		if signMultiplier*rev >= math.MaxInt32 || signMultiplier*rev <= math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return signMultiplier * rev
}
