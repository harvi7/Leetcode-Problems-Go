package twopointers

import "sort"

func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, low, high := -nums[i], i+1, len(nums)-1
		for low < high {
			sum := nums[low] + nums[high]
			if sum == target {
				results = append(results, []int{nums[i], nums[low], nums[high]})
				low++
				high--
				for low < high && nums[low] == nums[low-1] {
					low++
				}
				for low < high && nums[high] == nums[high+1] {
					high--
				}
			} else if sum > target {
				high--
			} else if sum < target {
				low++
			}
		}
	}
	return results
}
