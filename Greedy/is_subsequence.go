package greedy

func isSubsequence(s string, t string) bool {
	var LEFT_BOUND, RIGHT_BOUND = len(s), len(t)

	var p_left, p_right = 0, 0
	for p_left < LEFT_BOUND && p_right < RIGHT_BOUND {
		if s[p_left] == t[p_right] {
			p_left += 1
		}
		p_right += 1
	}
	return p_left == LEFT_BOUND
}
