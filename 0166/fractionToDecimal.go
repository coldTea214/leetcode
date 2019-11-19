import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	isNegative := false
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		isNegative = true
	}

	res := doFractionToDecimal(abs(numerator), abs(denominator))
	if isNegative {
		return "-" + res
	}
	return res
}

func doFractionToDecimal(numerator, denominator int) string {
	if numerator >= denominator {
		decimalPart := doFractionToDecimal(numerator%denominator, denominator)
		return strconv.Itoa(numerator/denominator) + decimalPart[1:]
	}

	digits := []byte{'0', '.'}
	numeratorAppearedInLoc := make(map[int]int)
	for i := 2; numerator != 0; i++ {
		if loc, ok := numeratorAppearedInLoc[numerator]; ok {
			return fmt.Sprintf("%s(%s)", string(digits[:loc]), string(digits[loc:]))
		}
		numeratorAppearedInLoc[numerator] = i

		numerator *= 10
		digits = append(digits, byte(numerator/denominator)+'0')
		numerator %= denominator
	}
	return string(digits)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
