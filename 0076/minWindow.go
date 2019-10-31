func minWindow(s string, t string) string {
	var letterInSWinCount, letterInTCount [128]int
	for i := range t {
		letterInTCount[t[i]]++
	}

	minWin := len(s) + 1
	var res string
	for i, j, count := 0, 0, 0; j < len(s); j++ {
		if letterInSWinCount[s[j]] < letterInTCount[s[j]] {
			// 出现了 window 中缺失的字母
			count++
		}
		letterInSWinCount[s[j]]++

		// 保证 window 不丢失所需字母的前提下
		// 让 i 尽可能的大
		for i <= j && letterInSWinCount[s[i]] > letterInTCount[s[i]] {
			letterInSWinCount[s[i]]--
			i++
		}

		curWin := j - i + 1
		if count == len(t) && minWin > curWin {
			minWin = curWin
			res = s[i : j+1]
		}

	}
	return res
}
