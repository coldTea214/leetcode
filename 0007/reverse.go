import "math"

// 无限制时的写法
func reverse(x int) int {
	var result int
    for x != 0 {
        digit := x % 10
        x /= 10
        result = result*10 + digit
    }
    if result < math.MinInt32 || result > math.MaxInt32 {
        return 0
    }
    return result
}

// 增加限制：环境无法存储64位整数
func reverse2(x int) int {
	var result int
    for x != 0 {
        if result < math.MinInt32/10 || result > math.MaxInt32/10 {
            return 0
        }
        digit := x % 10
        x /= 10
        result = result*10 + digit
    }
    return result
}
