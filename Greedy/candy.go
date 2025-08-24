package greedy

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candies := make([]int, n)
	// Initialize all with 1
	for i := range candies {
		candies[i] = 1
	}

	// Left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	sum := candies[n-1]

	// Right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i] < candies[i+1]+1 {
				candies[i] = candies[i+1] + 1
			}
		}
		sum += candies[i]
	}

	return sum
}
