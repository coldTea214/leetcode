import "strings"

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	wordsInThisLine := []string{}
	widthOfWordsInThisLine := 0
	isLastLine := false

	for !isLastLine {
		words, wordsInThisLine, widthOfWordsInThisLine, isLastLine = split(words, maxWidth)
		res = append(res, combine(wordsInThisLine, widthOfWordsInThisLine, maxWidth, isLastLine))
	}

	return res
}

func split(words []string, maxWidth int) ([]string, []string, int, bool) {
	wordsInThisLine := make([]string, 1)
	wordsInThisLine[0] = words[0]
	widthOfWordsInThisLine := len(words[0])

	i := 1
	for ; i < len(words); i++ {
		if widthOfWordsInThisLine+len(wordsInThisLine)+len(words[i]) > maxWidth {
			break
		}
		wordsInThisLine = append(wordsInThisLine, words[i])
		widthOfWordsInThisLine += len(words[i])
	}

	return words[i:], wordsInThisLine, widthOfWordsInThisLine, i == len(words)
}

func combine(words []string, widthOfWordsInThisLine, maxWidth int, isLastLine bool) string {
	wordCount := len(words)
	if wordCount == 1 || isLastLine {
		return combineSpecial(words, maxWidth)
	}

	spaceCount := wordCount - 1
	spaces := makeSpaces(spaceCount, maxWidth-widthOfWordsInThisLine)

	res := ""
	for i, space := range spaces {
		res += words[i] + space
	}
	if wordCount > 1 {
		res += words[wordCount-1]
	}
	return res
}

func combineSpecial(words []string, maxWidth int) string {
	res := strings.Join(words, " ")
	for len(res) < maxWidth {
		res += " "
	}
	return res
}

func makeSpaces(Len, count int) []string {
	res := make([]string, Len)
	for i := 0; i < count; i++ {
		res[i%Len] += " "
	}
	return res
}
