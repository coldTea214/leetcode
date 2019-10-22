// 执行用时 : 4 ms , 在所有 golang 提交中击败了 94.78% 的用户
// 内存消耗 : 2.2 MB , 在所有 golang 提交中击败了 80.85% 的用户

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		curLen := max(len1, len2)
		if curLen > end-start {
			start = i - (curLen-1)/2
			end = i + curLen/2
		}
	}
	return string(s[start : end+1])
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
