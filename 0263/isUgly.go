func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	primeNums := []int{2, 3, 5}
	for _, pNum := range primeNums {
		for num%pNum == 0 {
			num /= pNum
		}
	}
	return num == 1
}
