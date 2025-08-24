package string

func longestCommonPrefix(strs []string) string {
	output := ""
	temp := ""

	if len(strs) == 0 {
		return temp
	}

	l := len(strs[0])
	temp = strs[0]

	for _, str := range strs {
		if len(str) < l {
			l = len(str)
			temp = str
		}
	}

	flag := 0

	for i := 0; i < len(temp); i++ {
		c := temp[i]
		flag = 0
		for _, str := range strs {
			if !(str[i] == c) {
				flag = 0
				break
			} else {
				flag = 1
			}
		}
		if flag == 1 {
			output = output + string(c)
		} else {
			break
		}
	}

	return output
}
