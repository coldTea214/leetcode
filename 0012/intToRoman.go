import "bytes"

type Pair struct {
	Number int
	Rome   string
}

var numToRome = []Pair{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	res := bytes.Buffer{}

	for _, pair := range numToRome {
		for num >= pair.Number {
			res.WriteString(pair.Rome)
			num -= pair.Number
		}
	}

	return res.String()
}

