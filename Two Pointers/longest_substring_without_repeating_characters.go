package twopointers

// https://leetcode.com/problems/longest-substring-without-repeating-characters/solutions/4239419/go-solution-great-explanation-and-full-description/?envType=study-plan-v2&envId=top-interview-150

func lengthOfLongestSubstring(s string) int {
	store := make(map[uint8]int)
	var left, right, result int

	for right < len(s) {
		var r = s[right]
		store[r] += 1
		for store[r] > 1 {
			l := s[left]
			store[l] -= 1
			left += 1
		}
		result = max(result, right-left+1)

		right += 1
	}
	return result
}
