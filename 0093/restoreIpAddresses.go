import "fmt"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}

	res := []string{}
	parts := make([]string, 4)
	generateIP(0, 0, s, parts, &res)
	return res
}

func generateIP(idx, begin int, s string, parts []string, res *[]string) {
	if idx == 3 {
		temp := s[begin:]
		if isValid(temp) {
			parts[3] = temp
			*res = append(*res, fmt.Sprintf("%s.%s.%s.%s", parts[0], parts[1], parts[2], parts[3]))
		}
		return
	}

	maxRemain := 3 * (3 - idx)
	for end := begin + 1; end < len(s); end++ {
		// 如果剩下的太多
		if end+maxRemain < len(s) {
			continue
		}

		if end-begin > 3 {
			break
		}

		temp := s[begin:end]
		if isValid(temp) {
			parts[idx] = temp
			generateIP(idx+1, end, s, parts, res)
		}
	}
}

func isValid(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	if len(s) < 3 {
		return true
	}

	// len(s) == 3
	switch s[0] {
	case '1':
		return true
	case '2':
		if '0' <= s[1] && s[1] <= '4' {
			return true
		}
		if s[1] == '5' && '0' <= s[2] && s[2] <= '5' {
			return true
		}
	}
	return false
}
