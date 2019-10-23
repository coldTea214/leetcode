var romeToNum = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := romeToNum[s[len(s)-1]]
	for i := 1; i < len(s); i++ {
		if romeToNum[s[i-1]] >= romeToNum[s[i]] {
			res += romeToNum[s[i-1]]
		} else {
			res -= romeToNum[s[i-1]]
		}
	}
	return res
}
