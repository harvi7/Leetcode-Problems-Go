package math

func tribonacci(n int) int {
	if n < 3 {
		if n > 0 {
			return 1
		}
		return 0
	}

	a, b, c := 0, 1, 1
	for i := 0; i < n-2; i++ {
		tmp := a + b + c
		a = b
		b = c
		c = tmp
	}
	return c
}
