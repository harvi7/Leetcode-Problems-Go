// https://www.youtube.com/watch?v=w6LcypDgC4w

package stack

import "bytes"

func removeDuplicates(s string, k int) string {
	type stackElement struct {
		value rune
		count int
	}
	stack := []stackElement{}
	for _, v := range s {
		if len(stack) == 0 || v != stack[len(stack)-1].value {
			stack = append(stack, stackElement{
				value: v,
				count: 1,
			})
		} else {
			stack[len(stack)-1].count++
			if stack[len(stack)-1].count == k {
				stack = stack[:len(stack)-1]
			}
		}
	}
	res := bytes.Buffer{}
	for _, v := range stack {
		res.Write(bytes.Repeat([]byte{byte(v.value)}, v.count))
	}
	return res.String()
}
