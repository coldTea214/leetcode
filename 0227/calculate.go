func calculate(s string) int {
	stack := make([]int, 0, len(s))
	num, op := 0, '+'
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				tmp := stack[len(stack)-1] * num
				stack = stack[:len(stack)-1]
				stack = append(stack, tmp)
			case '/':
				tmp := stack[len(stack)-1] / num
				stack = stack[:len(stack)-1]
				stack = append(stack, tmp)
			}
			op = int32(s[i])
			num = 0
		}
	}

	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}
