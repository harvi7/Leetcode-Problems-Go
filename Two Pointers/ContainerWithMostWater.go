package twopointers

func maxArea(height []int) int {
	res, l, r := 0, 0, len(height)-1
	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
