package string

import "strings"

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for i := length / 2; i >= 1; i-- {
		if length%i == 0 {
			numRepeats := length / i
			substring := s[:i]
			sb := strings.Builder{}
			for j := 0; j < numRepeats; j++ {
				sb.WriteString(substring)
			}
			if sb.String() == s {
				return true
			}
		}
	}
	return false
}
