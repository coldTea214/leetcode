// 后置题 0097
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	count := make([]int, 256)
	for i := 0; i < len(s1); i++ {
		count[s1[i]]++
		count[s2[i]]--
	}
	for i := range count {
		if count[i] != 0 {
			return false
		}
	}

	// 检查子字符串是否 scramble
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[0:i], s2[0:i]) &&
			isScramble(s1[i:], s2[i:]) {
			return true
		}

		if isScramble(s1[0:i], s2[len(s1)-i:]) &&
			isScramble(s1[i:], s2[0:len(s1)-i]) {
			return true
		}
	}
	return false
}
