package backtracking

func generatePalindromes(s string) []string {
	mapCount := make([]int, 128)
	st := make([]rune, len(s)/2)

	if !canPermutePalindrome(s, mapCount) {
		return nil
	}

	var mid rune
	k := 0
	for i := 0; i < len(mapCount); i++ {
		if mapCount[i]%2 == 1 {
			mid = rune(i)
		}
		for j := 0; j < mapCount[i]/2; j++ {
			st[k] = rune(i)
			k++
		}
	}

	// local set to store unique permutations
	resultSet := make(map[string]struct{})

	var permute func([]rune, int)
	permute = func(arr []rune, l int) {
		if l == len(arr) {
			left := string(arr)
			var middle string
			if mid != 0 {
				middle = string(mid)
			}
			right := reverse(left)
			resultSet[left+middle+right] = struct{}{}
			return
		}

		for i := l; i < len(arr); i++ {
			if arr[l] != arr[i] || l == i { // avoid duplicates
				arr[l], arr[i] = arr[i], arr[l]
				permute(arr, l+1)
				arr[l], arr[i] = arr[i], arr[l] // backtrack
			}
		}
	}

	permute(st, 0)

	// Convert map keys to slice
	res := make([]string, 0, len(resultSet))
	for str := range resultSet {
		res = append(res, str)
	}
	return res
}

func canPermutePalindrome(s string, mapCount []int) bool {
	count := 0
	for _, ch := range s {
		mapCount[ch]++
		if mapCount[ch]%2 == 0 {
			count--
		} else {
			count++
		}
	}
	return count <= 1
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
