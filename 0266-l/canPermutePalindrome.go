func canPermutePalindrome(s string) bool {
	existOdd := map[byte]bool{}
	for i := range s {
		if existOdd[s[i]] {
			delete(existOdd, s[i])
		} else {
			existOdd[s[i]] = true
		}
	}
	return len(existOdd) <= 1
}
