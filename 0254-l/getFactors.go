package main

import (
	"fmt"
	"math"
)

func getFactors(n int) [][]int {
	var solutions [][]int
	doGetFactors(n, 2, int(math.Sqrt(float64(n))), []int{}, &solutions)
	return solutions
}

// end 可以节省大量计算量
func doGetFactors(n, beg, end int, solution []int, solutions *[][]int) {
	if n == 1 && len(solution) > 1 {
		deepcopy := make([]int, len(solution))
		copy(deepcopy, solution)
		*solutions = append(*solutions, deepcopy)
		return
	}
	
	for i := beg; i <= end; i++ {
		if n%i != 0 {
			continue
		}
		doGetFactors(n/i, i, int(math.Sqrt(float64(n/i))), append(solution, i), solutions)
	}
	if len(solution) != 0 {
		doGetFactors(1, -1, -1, append(solution, n), solutions)
	}
}

func main() {
	// fmt.Println(getFactors(1))
	// fmt.Println(getFactors(12))
	fmt.Println(getFactors(6811815))
}
