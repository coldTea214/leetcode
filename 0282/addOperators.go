import "strconv"

func addOperators(num string, target int) []string {
	expressions := []string{}

	var dfs func(string, string, int, int)
	dfs = func(num, expression string, result, prevAdd int) {
		if len(num) == 0 {
			if result == target {
				expressions = append(expressions, expression)
			}
			return
		}

		var curNum, leftNum string
		for i := 1; i <= len(num); i++ {
			curNum = num[:i]
			if curNum[0] == '0' && len(curNum) > 1 {
				return
			}

			cNum, _ := strconv.Atoi(curNum)
			leftNum = num[i:]

			if len(expression) == 0 {
				dfs(leftNum, curNum, cNum, cNum)
			} else {
				// prevAdd 保证 乘法 运算的优先性，遇到成法就把上一步已经 add 的数再减掉重新计算
				dfs(leftNum, expression+"+"+curNum, result+cNum, cNum)
				dfs(leftNum, expression+"-"+curNum, result-cNum, -cNum)
				dfs(leftNum, expression+"*"+curNum, result-prevAdd+prevAdd*cNum, prevAdd*cNum)
			}
		}
	}

	dfs(num, "", 0, 0)

	return expressions
}
