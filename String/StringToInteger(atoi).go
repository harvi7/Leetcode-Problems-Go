import (
	"math"
	"strings"
)

func myAtoi(s string) int {
    s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	// check for sign
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	// check for leading zeros
	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
	}
	// check for non-numeric characters
	if len(s) == 0 {
		return 0
	}
	if s[0] < '0' || s[0] > '9' {
		return 0
	}
	// check for overflow
	result := 0
	for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		result = result * 10 + int(s[0] - '0')
		if result > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		s = s[1:]
	}
	return result * sign
}