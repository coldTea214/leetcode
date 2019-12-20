import "strings"

// 解法1
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 解法2
func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

// 解法3
func strStr(haystack string, needle string) int {
	// kmp
}
