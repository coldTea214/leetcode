func isOneEditDistance(s, t string) bool {
	m, n := len(s), len(t)
	if m < n {
		return isOneEditDistance(t, s)
	}
	if m-n > 1 {
		return false
	}

	foundDifference := false
	for i, ch := range t {
		if s[i] != byte(ch) {
			foundDifference = true
			if m == n {
				return s[i+1:] == t[i+1:]
			}
			return s[i+1:] == t[i:]
		}
	}
	return foundDifference || m-n == 1
}
