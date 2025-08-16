package string

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for i := length / 2; i >= 1; i-- {
		if length%i == 0 {
			numRepeats := length / i
			substring := s[:i]
			sb := ""
			for j := 0; j < numRepeats; j++ {
				sb += substring
			}
			if sb == s {
				return true
			}
		}
	}
	return false
}
