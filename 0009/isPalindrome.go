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
