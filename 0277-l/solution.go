// 证明:
// 询问 a know b, b know a，出现 Y/Y,Y/N,N/Y,N/N 4种局面，可以确定 a&b都不是名人 或 a|b不是名人
// 也就是至少2次调用就可以排除1个人,经过2n次后,最多剩1个疑似名人,再通过n次调用确认是否是名人
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		candidate := 0
		for other := 0; other < n; other++ {
			if candidate == other {
				continue
			}
			if knows(candidate, other) || !knows(other, candidate) {
				candidate = other
			}
		}
		for other := 0; other < n; other++ {
			if candidate == other {
				continue
			}
			if knows(candidate, other) || !knows(other, candidate) {
				return -1
			}
		}
		return candidate
	}
}
