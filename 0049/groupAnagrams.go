package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	sameWord := make(map[[26]int][]string)
	for _, str := range strs {
		var letterCount [26]int
		for i := range str {
			letterCount[str[i]-'a']++
		}
		sameWord[letterCount] = append(sameWord[letterCount], str)
	}

	var result [][]string
	for _, value := range sameWord {
		result = append(result, value)
	}
	return result
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
}
