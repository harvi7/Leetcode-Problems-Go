func reverseWords(s string) string {
    result := []byte(s)
	left := 0
	for i, b := range result {
		if b == ' ' {
			reverse(&result, left, i - 1)
			left = i + 1
		}
	}
	reverse(&result, left, len(result) - 1)
	return string(result)
}

func reverse(src *[]byte, from int, to int) {
	for from < to {
		(*src)[from], (*src)[to] = (*src)[to], (*src)[from]
		from++
		to--
	}
}