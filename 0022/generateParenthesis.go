func generateParenthesis(n int) []string {
	var res []string
	doGenerate(n, 0, 0, "", &res)
	return res
}

func doGenerate(n, lNum, rNum int, cur string, res *[]string) {
	if lNum == n && rNum == n {
		*res = append(*res, cur)
		return
	}

	if lNum < n {
		doGenerate(n, lNum+1, rNum, cur+"(", res)
	}
	if rNum < lNum {
		doGenerate(n, lNum, rNum+1, cur+")", res)
	}
}
