package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1s := conv(version1)
	v2s := conv(version2)

	if len(v1s) != len(v2s) {
		v1s, v2s = toSameLen(v1s, v2s)
	}

	for i := 0; i < len(v1s); i++ {
		if v1s[i] < v2s[i] {
			return -1
		} else if v1s[i] > v2s[i] {
			return 1
		}
	}
	return 0
}

func conv(version string) []int {
	vs := strings.Split(version, ".")
	res := make([]int, len(vs))

	for i, v := range vs {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func toSameLen(v1s, v2s []int) ([]int, []int) {
	if len(v1s) > len(v2s) {
		newV2s := make([]int, len(v1s))
		copy(newV2s, v2s)
		return v1s, newV2s
	}

	v2s, newV1s := toSameLen(v2s, v1s)
	return newV1s, v2s
}

func main() {
	fmt.Println(compareVersion("1.0", "1.0.1"))
}
