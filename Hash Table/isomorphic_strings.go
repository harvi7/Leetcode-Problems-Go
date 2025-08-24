package hashtable

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte) // reverse mapping to avoid duplicate values

	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]

		if mapped, exists := sToT[a]; exists {
			if mapped != b {
				return false
			}
		} else {
			if mapped, exists := tToS[b]; exists {
				if mapped != a {
					return false
				}
			}
			sToT[a] = b
			tToS[b] = a
		}
	}

	return true
}
