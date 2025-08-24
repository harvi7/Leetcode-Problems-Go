package string

func reverseWords(s string) string {
	result := []byte(s)
	left := 0
	for i, b := range result {
		if b == ' ' {
			reverseW(&result, left, i-1)
			left = i + 1
		}
	}
	reverseW(&result, left, len(result)-1)
	return string(result)
}

func reverseW(src *[]byte, from int, to int) {
	for from < to {
		(*src)[from], (*src)[to] = (*src)[to], (*src)[from]
		from++
		to--
	}
}
