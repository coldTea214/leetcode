func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return nil
	}

	var res []string
	strCount := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		str := s[i : i+10]
		if strCount[str] == 1 {
			// 如果仅用bool来判断exist，res中会有重复串
			res = append(res, str)
		}
		strCount[str]++
	}
	return res
}
