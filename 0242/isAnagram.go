func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	count := make(map[rune]int, len(sr))
	for i := range sr {
		count[sr[i]]++
		count[tr[i]]--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}
