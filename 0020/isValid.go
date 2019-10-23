func isValid(s string) bool {
	sLen := len(s)

	stack := make([]byte, sLen)
	top := 0

	for i := 0; i < sLen; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = ')'
			top++
		case '[':
			stack[top] = ']'
			top++
		case '{':
			stack[top] = '}'
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}

	return top == 0
}
