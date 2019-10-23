import "bytes"

// 执行用时 : 4 ms , 在所有 golang 提交中击败了 96.31% 的用户
// 内存消耗 : 4.1 MB , 在所有 golang 提交中击败了 88.72% 的用户

// 示例行数为3，有些特征看的不明显，行数为5时
// L/0         I/8
// E/1     E/7 S/9      G/15
// E/2   D/6   H/10   N/14
// T/3 O/5     I/11 I/13
// C/4         R/12
func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	// p pace 步距，p=8
	p := numRows*2 - 2

	// 处理第一行
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// 处理中间的行
	for r := 1; r <= numRows-2; r++ {
		// 添加r行的第一个字符
		res.WriteByte(s[r])

		// k=8,16,24
		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// 处理最后一行
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}
