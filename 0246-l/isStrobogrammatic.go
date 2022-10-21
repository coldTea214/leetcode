func isStrobogrammatic(num string) bool {
	l, r := 0, len(num)-1
	for l <= r {
		if num[l] == '6' && num[r] == '9' ||
			num[l] == '9' && num[r] == '6' ||
			num[l] == '0' && num[r] == '0' ||
			num[l] == '1' && num[r] == '1' ||
			num[l] == '8' && num[r] == '8' {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
