package array

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = nums[i-1] * res[i-1]
	}

	prod := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = res[i] * prod
		prod = prod * nums[i]
	}
	return res

}
