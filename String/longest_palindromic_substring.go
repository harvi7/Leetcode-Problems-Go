package string

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var result string
	var start int
	var end int
	for i := 0; i < len(s)-1; i++ {
		start1, end1 := expandAroundCenter(s, i, i+1)
		start2, end2 := expandAroundCenter(s, i, i)
		if end1-start1 > end2-start2 {
			start = start1
			end = end1
		} else {
			start = start2
			end = end2
		}

		if start < end && end-start+1 > len(result) {
			result = s[start : end+1]
		}
	}
	if len(result) == 0 {
		return string(s[0])
	}
	return result
}

func expandAroundCenter(s string, start, end int) (int, int) {
	for start >= 0 && end < len(s) {
		if s[start] != s[end] {
			return start + 1, end - 1
		}
		start--
		end++
	}

	return start + 1, end - 1
}
