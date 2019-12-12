import "strconv"

func getHint(secret string, guess string) string {
	var bull, cow int
	var sCount, gCount [10]int

	for i := 0; i < len(secret); i++ {
		ns := int(secret[i] - '0')
		ng := int(guess[i] - '0')

		if ns == ng {
			bull++
		}
		sCount[ns]++
		gCount[ng]++
	}
	for i := 0; i <= 9; i++ {
		cow += min(sCount[i], gCount[i])
	}
	cow -= bull

	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
