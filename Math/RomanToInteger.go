func romanToInt(s string) int {
	charMap := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0

	for i := 0; i < len(s) - 1; i++ {
			if charMap[s[i]] < charMap[s[i + 1]] {
					total -= charMap[s[i]]
			} else {
					total += charMap[s[i]]
			}
	}
	
	total += charMap[s[len(s) - 1]]
	return total
}