func findSubstring(s string, words []string) []int {
	var res []int

	sLen := len(s)
	wordNum := len(words)
	if sLen == 0 || wordNum == 0 || sLen < wordNum*len(words[0]) {
		return res
	}
	wordLen := len(words[0])

	// sameWordNum 记录 words 中每个单词总共出现的次数
	sameWordNum := make(map[string]int, wordNum)
	for _, word := range words {
		sameWordNum[word]++
	}

	for i := 0; i < wordLen; i++ {
		// count 记录当前已经匹配了多少个 word
		// sameWordCount[word] 记录针对某一个具体 word 匹配了多少次
		count := 0
		sameWordCount := make(map[string]int, wordNum)
		left, right := i, i

		for left <= sLen-wordNum*wordLen {
			word := s[right : right+wordLen]
			right += wordLen

			num, ok := sameWordNum[word]
			if !ok {
				count = 0
				for key := range sameWordCount {
					delete(sameWordCount, key)
				}
				left = right
				continue
			}

			if sameWordCount[word] == num {
				t := s[left : left+wordLen]
				for sameWordCount[word] == num {
					count--
					sameWordCount[t]--
					left += wordLen

					t = s[left : left+wordLen]
				}
			}

			count++
			sameWordCount[word]++
			if count == wordNum {
				res = append(res, left)

				t := s[left : left+wordLen]
				count--
				sameWordCount[t]--
				left += wordLen
			}
		}
	}

	return res
}
