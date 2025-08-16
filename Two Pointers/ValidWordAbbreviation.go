package twopointers

// https://www.youtube.com/watch?v=Sut-F029biM

func validWordAbbreviation(word string, abbr string) bool {
	wordPtr, abbrPtr := 0, 0
	for wordPtr < len(word) && abbrPtr < len(abbr) {
		if isDigit(abbr[abbrPtr]) {
			steps := 0
			if abbr[abbrPtr] == '0' {
				return false
			}
			for abbrPtr < len(abbr) && isDigit(abbr[abbrPtr]) {
				steps = steps*10 + int(abbr[abbrPtr]-'0')
				abbrPtr += 1
			}
			wordPtr += steps
		} else {
			if word[wordPtr] != abbr[abbrPtr] {
				return false
			}
			wordPtr += 1
			abbrPtr += 1
		}
	}
	return wordPtr == len(word) && abbrPtr == len(abbr)
}

func isDigit(a byte) bool {
	return a >= '0' && a <= '9'
}
