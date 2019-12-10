import "strings"

var lt21 = []string{
	"",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Eleven",
	"Twelve",
	"Thirteen",
	"Fourteen",
	"Fifteen",
	"Sixteen",
	"Seventeen",
	"Eighteen",
	"Nineteen",
	"Twenty",
}

var ten = []string{
	"",
	"",
	"Twenty",
	"Thirty",
	"Forty",
	"Fifty",
	"Sixty",
	"Seventy",
	"Eighty",
	"Ninety",
}

var thousand = []string{
	"",
	"Thousand",
	"Million",
	"Billion",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	res := ""
	for i := 0; num > 0; i++ {
		if num%1000 != 0 {
			res = numLt1000(num%1000) + thousand[i] + " " + res
		}
		num /= 1000
	}
	return strings.TrimRight(res, " ")
}

func numLt1000(num int) string {
	if num == 0 {
		return ""
	}
	if num <= 20 {
		return lt21[num] + " "
	}
	if num < 100 {
		return ten[num/10] + " " + numLt1000(num%10)
	}
	return lt21[num/100] + " Hundred " + numLt1000(num%100)
}
