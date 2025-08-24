package math

func fib(n int) int {
	if n <= 1 {
		return n
	}
	current, prev1, prev2 := 0, 1, 0

	for i := 2; i <= n; i++ {
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}
	return current
}
