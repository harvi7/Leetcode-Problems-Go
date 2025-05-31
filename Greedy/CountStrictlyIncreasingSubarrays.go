package greedy

// https://leetcode.com/problems/count-strictly-increasing-subarrays/solutions/

func countSubarrays(nums []int) int64 {
	var subarrayCount int64 = 0

	for i := 0; i < len(nums); i++ {
		currSubarray := 1
		for i+1 < len(nums) && nums[i] < nums[i+1] {
			currSubarray++
			i++
		}
		subarrayCount += int64((currSubarray * (currSubarray + 1)) / 2)
	}

	return subarrayCount
}
