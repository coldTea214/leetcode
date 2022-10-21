package main

import "fmt"

func canWin(currentState string) bool {
	return canWinHelper(make(map[string]bool), currentState)
}

// +++++ false
//   --+++ true
//     ----+ false
//   +--++ true
//     +---- false
//   ++--+ true
//     ----+ false # 重复
//   +++-- true
//     --+-- false
func canWinHelper(cache map[string]bool, state string) bool {
	if win, ok := cache[state]; ok {
		return win
	}
	for i := 0; i < len(state)-1; i++ {
		if state[i:i+2] == "++" {
			newState := state[:i] + "--" + state[i+2:]
			if !canWinHelper(cache, newState) {
				cache[state] = true
				return true
			}
		}
	}
	cache[state] = false
	return false
}

func main() {
	fmt.Println(canWin("++++++"))
}
