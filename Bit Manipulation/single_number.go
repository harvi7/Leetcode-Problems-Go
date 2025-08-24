// https://leetcode.com/problems/single-number/editorial/?envType=study-plan-v2&envId=top-interview-150

package bitManipulation

func singleNumber(nums []int) int {
	// result := -1
	// m := make(map[int]int, len(nums))
	// for _, v := range nums {
	// 	m[v]++
	// }
	// for key, valueMap := range m {
	// 	if valueMap == 1 {
	// 		result = key
	// 	}
	// }
	// return result

	// set := make(map[int]bool)
	// sumOfSet, sumOfNums := 0, 0
	// for _, num := range nums {
	//     if !set[num] {
	//         set[num] = true
	//         sumOfSet += num
	//     }
	//     sumOfNums += num
	// }
	// return 2*sumOfSet - sumOfNums

	a := 0
	for _, i := range nums {
		a ^= i
	}
	return a
}
