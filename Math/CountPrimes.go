package math

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	res := 0
	isPrime := make([]bool, n+1, n+1)
	for i, _ := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			res++
			for j := i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return res
}
