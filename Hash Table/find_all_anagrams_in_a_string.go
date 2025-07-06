package hashtable

func findAnagrams(s string, p string) []int {
	ns, np := len(s), len(p)
	if ns < np {
		return []int{}
	}

	pCount := make([]int, 26)
	sCount := make([]int, 26)

	for _, ch := range p {
		pCount[ch-'a']++
	}

	output := []int{}
	for i := 0; i < ns; i++ {
		sCount[s[i]-'a']++
		if i >= np {
			sCount[s[i-np]-'a']--
		}
		if equal(pCount, sCount) {
			output = append(output, i-np+1)
		}
	}
	return output
}

func equal(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
