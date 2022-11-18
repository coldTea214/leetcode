package main

import (
	"fmt"
	"strings"
)

var pairs = [][2]byte{{'0', '0'}, {'1', '1'}, {'6', '9'}, {'8', '8'}, {'9', '6'}}
var base = []string{"0", "1", "8"}

func findStrobogrammatic(n int) []string {
	var strobogrammatics []string
	if n%2 == 1 {
		for _, s := range base {
			findStrobogrammaticHelper(n-1, s, &strobogrammatics)
		}
	} else {
		findStrobogrammaticHelper(n, "", &strobogrammatics)
	}
	return strobogrammatics
}

func findStrobogrammaticHelper(i int, strobogrammatic string, strobogrammatics *[]string) {
	if i == 0 {
		*strobogrammatics = append(*strobogrammatics, strobogrammatic)
		return
	}

	for _, pair := range pairs {
		if i == 2 && pair[0] == '0' {
			continue
		}
		num := &strings.Builder{}
		num.WriteByte(pair[0])
		num.WriteString(strobogrammatic)
		num.WriteByte(pair[1])
		findStrobogrammaticHelper(i-2, num.String(), strobogrammatics)
	}
}

func main() {
	fmt.Println(findStrobogrammatic(1))
	fmt.Println(findStrobogrammatic(2))
	fmt.Println(findStrobogrammatic(3))
	fmt.Println(findStrobogrammatic(4))
}
