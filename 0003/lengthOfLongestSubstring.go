func lengthOfLongestSubstring(s string) int {
	letterIdx := make(map[byte]int)
	strBeg, maxLen := 0, 0
	for idx1 := 0; idx1 < len(s); idx1++ {
		letter := s[idx1]

		idx2, ok := letterIdx[letter]
		if ok {
			if idx1-strBeg > maxLen {
				maxLen = idx1 - strBeg
			}

			for i := strBeg; i <= idx2; i++ {
				delete(letterIdx, s[i])
			}
			strBeg = idx2 + 1
		}

		letterIdx[letter] = idx1
	}
	if len(s)-strBeg > maxLen {
		maxLen = len(s) - strBeg
	}
	return maxLen
}
