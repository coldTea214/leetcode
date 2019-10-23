func lengthOfLongestSubstring(s string) int {
	letterIdx := make(map[byte]int)
	strBeg, maxLen, curLen := 0, 0, 0
	for idx1 := 0; idx1 < len(s); idx1++ {
		letter := s[idx1]

		idx2, ok := letterIdx[letter]
		if ok {
			curLen = idx1 - strBeg
			if curLen > maxLen {
				maxLen = curLen
			}

			for i := strBeg; i <= idx2; i++ {
				delete(letterIdx, s[i])
			}
			strBeg = idx2 + 1
		}

		letterIdx[letter] = idx1
	}
	curLen = len(s) - strBeg
	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}
