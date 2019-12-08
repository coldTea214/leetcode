import "strconv"

func diffWaysToCompute(input string) []int {
	// 缓存已经计算过的字符串
	calculated := make(map[string][]int)
	var dfs func(string) []int
	dfs = func(s string) []int {
		res := make([]int, 0, len(s))
		if t, ok := calculated[s]; ok {
			return t
		}

		for i := range s {
			if s[i] == '+' || s[i] == '-' || s[i] == '*' {
				// 此时，s[i] 作为最后一个运算的运算符
				for _, left := range dfs(s[:i]) {
					for _, right := range dfs(s[i+1:]) {
						res = append(res, operate(left, right, s[i]))
					}
				}
			}
		}

		// s 中不存在运算符
		if len(res) == 0 {
			temp, _ := strconv.Atoi(s)
			res = append(res, temp)
		}

		calculated[s] = res
		return res
	}

	return dfs(input)
}

func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return a * b
	}
}
