package main

import "fmt"

func generatePossibleNextMoves(currentState string) []string {
	if len(currentState) <= 1 {
		return nil
	}

	var ans []string
	state := []byte(currentState)
	for i := 0; i < len(state)-1; i++ {
		if state[i] == '+' && state[i+1] == '+' {
			state[i], state[i+1] = '-', '-'
			ans = append(ans, string(state))
			state[i], state[i+1] = '+', '+'
		}
	}
	return ans
}

func main() {
	fmt.Println(generatePossibleNextMoves("++++"))
}
