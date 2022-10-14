package main

import (
	"fmt"
	"strconv"
)

func addOperators(num string, target int) []string {
	solutions := []string{}
	doAddOperators(num, target, 0, 0, "", &solutions)
	return solutions
}

func doAddOperators(num string, target int, result, prevAdd int, solution string, solutions *[]string) {
	if len(num) == 0 {
		if result == target {
			*solutions = append(*solutions, solution)
		}
		return
	}

	for i := 1; i <= len(num); i++ {
		curNum := num[:i]
		if curNum[0] == '0' && len(curNum) > 1 {
			return
		}

		cNum, _ := strconv.Atoi(curNum)
		if len(solution) == 0 {
			doAddOperators(num[i:], target, cNum, cNum, curNum, solutions)
		} else {
			// prevAdd 保证 乘法 运算的优先性，遇到成法就把上一步已经 add 的数再减掉重新计算
			doAddOperators(num[i:], target, result+cNum, cNum, solution+"+"+curNum, solutions)
			doAddOperators(num[i:], target, result-cNum, -cNum, solution+"-"+curNum, solutions)
			doAddOperators(num[i:], target, result-prevAdd+prevAdd*cNum, prevAdd*cNum, solution+"*"+curNum, solutions)
		}
	}
}

func main() {
	// [1+2+3 1*2*3]
	fmt.Println(addOperators("123", 6))
}
