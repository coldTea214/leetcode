package main

import "fmt"

// 后置题 0131
// manacher 算法更快
func longestPalindrome(s string) string {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := range s {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true
		if i > 0 {
			isPalindrome[i][i-1] = true // 无物理含义, 方便下面dp迭代
		}
	}

	start, end := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && isPalindrome[i+1][j-1] {
				isPalindrome[i][j] = true
				if j-i > end-start {
					start, end = i, j
				}
			}
		}
	}
	return s[start : end+1]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
