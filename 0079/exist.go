package main

import "fmt"

// dfs
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existHelper(i, j, word, board, visited) {
				return true
			}
		}
	}
	return false
}

func existHelper(i, j int, word string, board [][]byte, visited [][]bool) bool {
	if word == "" {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		visited[i][j] ||
		board[i][j] != word[0] {
		return false
	}

	visited[i][j] = true

	if existHelper(i-1, j, word[1:], board, visited) ||
		existHelper(i+1, j, word[1:], board, visited) ||
		existHelper(i, j-1, word[1:], board, visited) ||
		existHelper(i, j+1, word[1:], board, visited) {
		return true
	}

	visited[i][j] = false

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
}
