package main

import "fmt"

func convertToTitle(n int) string {
	res := ""

	for n > 0 {
		res = string(byte((n-1)%26)+'A') + res
		n = (n - 1) / 26
	}

	return res
}

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
}
