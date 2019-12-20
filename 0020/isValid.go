func isValid(s string) bool {
	stack := make([]byte, len(s))
	top := 0

	for i := 0; i < len(s); i++ {
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
