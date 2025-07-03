package hashtable

import "strconv"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	ans := make(map[string][]string)

	for _, s := range strs {
		count := make([]int, 26)
		for _, c := range s {
			count[c-'a']++
		}
		key := ""
		for i := 0; i < 26; i++ {
			key += "#"
			key += strconv.Itoa(count[i])
		}
		if _, ok := ans[key]; !ok {
			ans[key] = []string{}
		}
		ans[key] = append(ans[key], s)
	}
	result := [][]string{}
	for _, v := range ans {
		result = append(result, v)
	}
	return result
}
