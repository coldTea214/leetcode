// 前置题 read 仅调用一次，这题会调用多次
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	buf4 := make([]byte, 4)
	l, r := 0, 0
	return func(buf []byte, n int) int {
		if len(buf) > n {
			buf = buf[:n]
		}
		next := 0
		if l > 0 {
			// 上一轮调用 read4 还有未读完字符
			count := copy(buf, buf4[l:r])
			l = (l + count) % r
			if l > 0 {
				return count
			} else {
				next = count
			}
		}
		for next < len(buf) {
			r = read4(buf4)
			if r == 0 {
				return next
			}
			// buf = buf[:n] 确保了不会 copy 超过 n 个字符
			count := copy(buf[next:], buf4[:r])
			next += count
			l = (l + count) % r
		}
		return next
	}
}
