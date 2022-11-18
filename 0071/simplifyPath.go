import "strings"

func simplifyPath(path string) string {
	pLen := len(path)
	stack := make([]string, 0, pLen/2)
	dir := make([]byte, 0, pLen)

	for i := 0; i < pLen; i++ {
		// 使用前，清空 dir
		// 这个方法比 dir = []byte{} 快了近 8 倍
		dir = dir[:0]
		for i < pLen && path[i] != '/' {
			dir = append(dir, path[i])
			i++
		}

		s := string(dir)
		switch s {
		case ".", "":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/")
}
