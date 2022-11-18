package main

import "fmt"

// dfs
var numberToLetter = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var words []string
	generateWords(digits, "", &words)
	return words
}

// 只要用了append, 就需要使用slice指针 
func generateWords(digits, word string, words *[]string) {
	if digits == "" {
		if word != "" {
			*words = append(*words, word)
		}
		return
	}

	for _, letter := range numberToLetter[digits[0]] {
		generateWords(digits[1:], word+letter, words)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
