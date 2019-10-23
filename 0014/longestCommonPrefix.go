func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	var prefix []byte
	for j := 0; j < minLen; j++ {
		letter := strs[0][j]

		for i := 1; i < len(strs); i++ {
			if strs[i][j] != letter {
				return string(prefix)
			}
		}

		prefix = append(prefix, letter)
	}
	return string(prefix)
}
