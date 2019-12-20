func countAndSay(n int) string {
	buf := []byte{'1'}
	for ; n > 1; n-- {
		buf = say(buf)
	}
	return string(buf)
}

func say(buf []byte) []byte {
	var res []byte
	for i, j := 0, 1; i < len(buf); {
		for j < len(buf) && buf[j] == buf[i] {
			j++
		}

		res = append(res, byte(j-i+'0'), buf[i])
		i = j
	}
	return res
}
