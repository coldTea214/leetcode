func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	i := getIndex(s + "#" + reverse(s))
	return reverse(s[i:]) + s
}

func reverse(s string) string {
	bytes := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		bytes[i] = s[len(s)-1-i]
	}
	return string(bytes)
}

func getIndex(s string) int {
	// table[i] 是 s[:i+1] 的前缀集与后缀集中，最长公共元素的长度
	// "abcd" 的前缀集是 ["", "a", "ab", "abc"]
	// "abcd" 的后缀集是 ["", "d", "cd", "bcd"]
	// "abcd" 的前缀集与后缀集的最长公共元素是"", 它的长度是 0
	table := make([]int, len(s))
	klen, i := 0, 1

	for i < len(s) {
		if s[i] == s[klen] {
			klen++
			table[i] = klen
			i++
		} else {
			if klen == 0 {
				table[i] = 0
				i++
			} else {
				klen = table[klen-1]
			}
		}
	}

	return table[len(s)-1]
}
