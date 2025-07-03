package greedy

// func jump(nums []int) int {
//     var (
//         currEnd int
//         farthest int
//         jumps int
//     )

//     for i:=0;i<len(nums) - 1;i++ {

//         if (farthest < i + nums[i]) {
//             farthest = i + nums[i]
//         }
//         if (i == currEnd) {
//             jumps++
//             currEnd = farthest
//         }
//     }
//     return jumps
// }

func jump(nums []int) int {
	res, l, r, farthest := 0, 0, 0, 0

	for r < len(nums)-1 {
		for i := l; i < r+1; i++ {
			farthest = max(farthest, i+nums[i])
		}
		l = r + 1
		r = farthest
		res += 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
