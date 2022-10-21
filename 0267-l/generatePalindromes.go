package main

import "fmt"

func generatePalindromes(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	count := make(map[byte]int)
	for i := range s {
		count[s[i]]++
	}
	var halfLetters []byte
	oddLetter, oddCnt := "", 0
	for k, v := range count {
		if v%2 != 0 {
			oddLetter = string(k)
			oddCnt++
			if oddCnt > 1 {
				return nil
			}
		}
		count[k] /= 2
		for i := 0; i < count[k]; i++ {
			halfLetters = append(halfLetters, k)
		}
	}
	var solutions [][]byte
	doGenerateWords(halfLetters, make(map[int]bool), []byte{}, &solutions)
	ans := make([]string, len(solutions))
	for i, word := range solutions {
		ans[i] = string(word) + oddLetter
		ans[i] += string(reverse(word))
	}
	return ans
}

func doGenerateWords(halfLetters []byte, letterInWord map[int]bool, solution []byte, solutions *[][]byte) {
	if len(solution) == len(halfLetters) {
		deepcopy := make([]byte, len(halfLetters))
		copy(deepcopy, solution)
		*solutions = append(*solutions, deepcopy)
		return
	}

	letterInWordThisLevel := make(map[byte]bool)
	for i := 0; i < len(halfLetters); i++ {
		if !letterInWord[i] && !letterInWordThisLevel[halfLetters[i]] {
			letterInWord[i] = true
			letterInWordThisLevel[halfLetters[i]] = true
			doGenerateWords(halfLetters, letterInWord, append(solution, halfLetters[i]), solutions)
			letterInWord[i] = false
		}
	}
}

func reverse(str []byte) []byte {
	left, right := 0, len(str)-1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	return str
}

func main() {
	// fmt.Println(generatePalindromes("aaa"))
	// fmt.Println(generatePalindromes("aabb"))
	// fmt.Println(generatePalindromes("aabbcc"))
	fmt.Println(generatePalindromes("aaaaaa"))
}
