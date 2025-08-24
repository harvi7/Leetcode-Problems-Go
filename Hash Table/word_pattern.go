package hashtable

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte) // to handle reverse mapping

	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		w := words[i]

		if mappedWord, exists := charToWord[c]; exists {
			if mappedWord != w {
				return false
			}
		} else {
			if mappedChar, exists := wordToChar[w]; exists {
				if mappedChar != c {
					return false
				}
			}
			charToWord[c] = w
			wordToChar[w] = c
		}
	}

	return true
}
