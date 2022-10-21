package main

import "fmt"

func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				wallsAndGatesHelper(rooms, i, j, 0)
			}
		}
	}
}

func wallsAndGatesHelper(rooms [][]int, i, j, step int) {
	if i < 0 || i >= len(rooms) || j < 0 || j >= len(rooms[0]) || rooms[i][j] < step {
		return
	}

	rooms[i][j] = step
	wallsAndGatesHelper(rooms, i-1, j, step+1)
	wallsAndGatesHelper(rooms, i+1, j, step+1)
	wallsAndGatesHelper(rooms, i, j-1, step+1)
	wallsAndGatesHelper(rooms, i, j+1, step+1)
}

func main() {
	rooms := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}
	wallsAndGates(rooms)
	// [[3 -1 0 1] [2 2 1 -1] [1 -1 2 -1] [0 -1 3 4]]
	fmt.Println(rooms)
}
