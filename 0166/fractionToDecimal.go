package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	var res string
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		res = "-"
	}
	numerator, denominator = abs(numerator), abs(denominator)

	// 先处理整数部分
	res += strconv.Itoa(numerator / denominator)
	numerator %= denominator

	// 再处理小数部分
	digits := []byte{'.'}
	numeratorAppearedInLoc := make(map[int]int)
	for i := 1; numerator != 0; i++ {
		if _, ok := numeratorAppearedInLoc[numerator]; ok {
			break
		}
		numeratorAppearedInLoc[numerator] = i

		numerator *= 10
		digits = append(digits, byte(numerator/denominator)+'0')
		numerator %= denominator
	}

	if numerator == 0 {
		res += string(digits)
	} else {
		// 有循环小数
		loc := numeratorAppearedInLoc[numerator]
		res += fmt.Sprintf("%s(%s)", string(digits[:loc]), string(digits[loc:]))
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// fmt.Println(fractionToDecimal(337, 333))
	fmt.Println(fractionToDecimal(-50, 8))
}
