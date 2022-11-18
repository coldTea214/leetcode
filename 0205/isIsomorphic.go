func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sLastAppearLoc := make([]int, 256)
	tLastAppearLoc := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if sLastAppearLoc[int(s[i])] != tLastAppearLoc[int(t[i])] {
			return false
		}

		sLastAppearLoc[int(s[i])] = i + 1
		tLastAppearLoc[int(t[i])] = i + 1
	}

	return true
}
