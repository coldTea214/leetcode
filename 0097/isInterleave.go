func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return doIsInterleave(s1, s2, s3)
}

func doIsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}

	if s1[0] == s3[0] {
		if doIsInterleave(s1[1:], s2, s3[1:]) {
			return true
		}
	}
	if s2[0] == s3[0] {
		return doIsInterleave(s1, s2[1:], s3[1:])
	}
	return false
}

