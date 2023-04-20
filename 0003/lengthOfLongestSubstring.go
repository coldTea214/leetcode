func lengthOfLongestSubstring(s string) int {
	letterIdx := make(map[byte]int)
	strBeg, maxLen := 0, 0
	for idx1 := 0; idx1 < len(s); idx1++ {
		// 直接 range s，获得的字母是 rune 类型，这种方式是 byte 类型
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
