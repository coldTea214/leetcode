import "fmt"

func groupAnagrams(strs []string) [][]string {
	sameWord := make(map[string][]string)
	for _, str := range strs {
		var letterCount [26]int
		for i := range str {
			letterCount[str[i]-'a']++
		}

		var countStr string
		for _, count := range letterCount {
			countStr += fmt.Sprintf("%v/", count)
		}

		sameWord[countStr] = append(sameWord[countStr], str)
	}

	var result [][]string
	for _, value := range sameWord {
		result = append(result, value)
	}
	return result
}
