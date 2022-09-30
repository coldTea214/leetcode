// 更常见的应该还是 manacher 算法

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// 不用 # 分隔，i 会超过原本字符串长度
	i := kmp(s + "#" + reverse(s))
	return reverse(s[i:]) + s
}

func reverse(s string) string {
	bytes := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		bytes[i] = s[len(s)-1-i]
	}
	return string(bytes)
}

// kmp 算法，介绍可参考：http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
// 输入: acecaaa
// 构造: acecaaa#aaaceca
// 跳表: 000011101112345
// 最后一个值为 5 含义是：后缀 aceca 跟前缀一致
func kmp(s string) int {
	next := make([]int, len(s))
	klen, i := 0, 1

	for i < len(s) {
		if s[i] == s[klen] {
			klen++
			next[i] = klen
			i++
		} else {
			if klen == 0 {
				next[i] = 0
				i++
			} else {
				klen = next[klen-1]
			}
		}
	}

	return next[len(s)-1]
}
