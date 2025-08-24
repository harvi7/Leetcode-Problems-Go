// https://leetcode.com/problems/group-shifted-strings/editorial/

package string

func shiftLetter(letter byte, shift byte) byte {
	return (letter-shift+26)%26 + 'a'
}

func getHash(s string) string {
	chars := []byte(s)

	// number of shifts to make first char 'a'
	shift := chars[0]
	for i := 0; i < len(chars); i++ {
		chars[i] = shiftLetter(chars[i], shift)
	}
	return string(chars)
}

func groupStrings(strings []string) [][]string {
	mapHashToList := make(map[string][]string)

	// create hash key for each string
	for _, str := range strings {
		hashKey := getHash(str)
		mapHashToList[hashKey] = append(mapHashToList[hashKey], str)
	}
	// collect groups
	groups := make([][]string, 0, len(mapHashToList))
	for _, group := range mapHashToList {
		groups = append(groups, group)
	}

	return groups
}
