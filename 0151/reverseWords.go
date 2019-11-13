import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}
	return strings.Join(words, " ")
}
