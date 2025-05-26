package string

import "bytes"

func licenseKeyFormatting(s string, k int) string {
	var b bytes.Buffer
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch == '-' {
			continue
		} else if ch >= 'a' && ch <= 'z' {
			ch = byte(int(ch) + (int('Z') - int('z')))
		}
		b.WriteByte(ch)
		count++
		if count == k {
			count = 0
			b.WriteByte('-')
		}
	}

	result := b.Bytes()
	if len(result) > 0 && result[len(result)-1] == '-' {
		result = result[:len(result)-1]
	}
	return reverse(result)
}

func reverse(b []byte) string {
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	return string(b)
}
