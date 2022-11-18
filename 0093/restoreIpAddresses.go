package main

import "fmt"

// dfs
func restoreIpAddresses(s string) []string {
	solutions := []string{}
	restoreIpAddressesHelper(s, []string{}, &solutions)
	return solutions
}

func restoreIpAddressesHelper(s string, solution []string, solutions *[]string) {
	if len(solution) == 3 {
		if isValid(s) {
			solution = append(solution, s)
			*solutions = append(*solutions, fmt.Sprintf("%s.%s.%s.%s", solution[0], solution[1], solution[2], solution[3]))
		}
		return
	}

	t := 4 - len(solution)
	if len(s) < t || len(s) > 3*t {
		return
	}

	for end := 1; end <= 3 && end < len(s); end++ {
		if isValid(s[:end]) {
			restoreIpAddressesHelper(s[end:], append(solution, s[:end]), solutions)
		}
	}
}

func isValid(s string) bool {
	if len(s) > 3 {
		return false
	}
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

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
