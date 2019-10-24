// ) ( ( ) ( ) ) ) ( ( ( ( ( ) ) ) ) (
// 第一遍顺序访问
// 访问到（，left++，访问到），left--，只要left>=0，记录当前以）为结尾存在有效括号串
// 0 0 0 1 0 1 1 0 0 0 0 0 0 1 1 1 1 0
// 第二遍逆序访问
// 确认这些有效括号串的起始点在哪
// 0 1 1 1 1 1 1 0 0 1 1 1 1 1 1 1 1 0
// 第三遍再顺序访问
// 主要是看下这些独立的有效串能不能连起来

func longestValidParentheses(s string) int {
	record := make([]bool, len(s))
	var left, right int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if left < 0 {
				left = 0
			}
			left++
		} else {
			left--
			if left >= 0 {
				record[i] = true
			}
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			if right < 0 {
				right = 0
			}
			right++
		} else {
			right--
			if right >= 0 {
				record[i] = true
			}
		}
	}

	var count, max int
	for i := 0; i < len(s); i++ {
		if record[i] {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		max = count
	}
	return max
}
