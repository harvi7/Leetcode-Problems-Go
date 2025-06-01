package array

// https://leetcode.com/problems/fair-candy-swap/editorial/

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sa, sb := 0, 0
	m := map[int]bool{} // set of aliceSizes
	for _, val := range aliceSizes {
		sa += val
		m[val] = true
	}
	for _, val := range bobSizes {
		sb += val
	}

	for _, val := range bobSizes {
		neededAlice := (sa-sb)/2 + val // we try to exchange Bob's val with Alice's neededAlice
		if _, ok := m[neededAlice]; ok {
			return []int{neededAlice, val}
		}
	}
	return []int{0, 0}
}
