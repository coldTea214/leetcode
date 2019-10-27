func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	var res int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}
	return res
}
