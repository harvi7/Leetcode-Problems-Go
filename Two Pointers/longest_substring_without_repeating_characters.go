func lengthOfLongestSubstring(s string) int {
    m, ans, i := make(map[rune]int), 0, 0
	for j, c := range s {
		if _, okay := m[c]; okay == true && m[c] >= i {
			if j - i > ans {
				ans = j - i
			}
			i = m[c] + 1
		}
		m[c] = j
	}
	if len(s) - i > ans {
		ans = len(s) - i
	}
	return ans
}