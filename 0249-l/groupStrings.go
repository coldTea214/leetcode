package main

import (
	"fmt"
	"strings"
)

func groupStrings(strings []string) [][]string {
	ansMap := make(map[string][]string)
	for _, str := range strings {
		key := genKey(str)
		ansMap[key] = append(ansMap[key], str)
	}
	ans := [][]string{}
	for _, v := range ansMap {
		ans = append(ans, v)
	}
	return ans
}

func genKey(v string) string {
	key := &strings.Builder{}
	for i := 0; i < len(v); i++ {
		t := (v[i] - v[0] + 26) % 26
		key.WriteByte(t)
		key.WriteByte('.')
	}
	return key.String()
}

func main() {
	input := []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}
	fmt.Println(groupStrings(input))
}
