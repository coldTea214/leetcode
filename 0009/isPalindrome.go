// 执行用时 : 20 ms , 在所有 golang 提交中击败了 67.22% 的用户
// 内存消耗 : 5 MB , 在所有 golang 提交中击败了 98.63% 的用户

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	number := x
	var reverseNumber int
	for x != 0 {
		reverseNumber = reverseNumber*10 + x%10
		x /= 10
	}
	return reverseNumber == number
}
