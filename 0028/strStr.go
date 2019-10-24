// 从题目定义为‘简单’，估摸着也不是想让实现kmp了

func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}

	return -1
}
