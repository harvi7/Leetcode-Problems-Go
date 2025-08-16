package twopointers

// func rotate(nums []int, k int)  {
//     if len(nums) == 1 {
// 		return
// 	}
// 	k = k % len(nums)
// 	newline := append(nums[len(nums) - k:], nums[:len(nums) - k]...)
// 	copy(nums, newline)
// }

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverseRotate(nums, 0, n-1)
	reverseRotate(nums, 0, k-1)
	reverseRotate(nums, k, n-1)
}

func reverseRotate(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
