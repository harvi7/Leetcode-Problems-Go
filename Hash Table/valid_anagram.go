package hashtable

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	char_counts := make([]int, 26)

	for i := 0; i < len(s); i++ {
		char_counts[s[i]-'a']++
		char_counts[t[i]-'a']--
	}

	for _, count := range char_counts {
		if count != 0 {
			return false
		}
	}
	return true
}
