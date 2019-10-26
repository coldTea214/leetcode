func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	numByte1 := []byte(num1)
	numByte2 := []byte(num2)
	resInt := make([]int, len(num1)+len(num2))

	for i := 0; i < len(numByte1); i++ {
		for j := 0; j < len(numByte2); j++ {
			resInt[i+j+1] += int(numByte1[i]-'0') * int(numByte2[j]-'0')
		}
	}

	for i := len(resInt) - 1; i > 0; i-- {
		resInt[i-1] += resInt[i] / 10
		resInt[i] = resInt[i] % 10
	}

	if resInt[0] == 0 {
		resInt = resInt[1:]
	}

	resByte := make([]byte, len(resInt))
	for i := 0; i < len(resInt); i++ {
		resByte[i] = '0' + byte(resInt[i])
	}

	return string(resByte)
}

