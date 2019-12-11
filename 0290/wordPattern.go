import "strings"

func wordPattern(pattern string, str string) bool {
	ps := strings.Split(pattern, "")
	ss := strings.Split(str, " ")
	if len(ps) != len(ss) {
		return false
	}
	return isMatch(ps, ss) && isMatch(ss, ps)
}

func isMatch(strs1, strs2 []string) bool {
	relation := make(map[string]string, len(strs1))
	for i := 0; i < len(strs1); i++ {
		word1, word2 := strs1[i], strs2[i]
		if relation[word1] == "" {
			relation[word1] = word2
		} else if relation[word1] != word2 {
			return false
		}
	}
	return true
}
