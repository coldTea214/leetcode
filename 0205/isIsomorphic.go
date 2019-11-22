func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sLastAppearInLoc := make([]int, 256)
	tLastAppearInLoc := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if sLastAppearInLoc[int(s[i])] != tLastAppearInLoc[int(t[i])] {
			return false
		}

		sLastAppearInLoc[int(s[i])] = i + 1
		tLastAppearInLoc[int(t[i])] = i + 1
	}

	return true
}
