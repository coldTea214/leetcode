// 并查集
func validTree(n int, edges [][]int) bool {
	father := make([]int, n)
	for i := 0; i < n; i++ {
		father[i] = i
	}

	for _, edge := range edges {
		root1 := search(father, edge[0])
		root2 := search(father, edge[1])
		if root1 == root2 {
			return false
		}

		// 反过来也没区别
		father[root2] = root1
	}

	root := search(father, 0)
	for i := 1; i < n; i++ {
		if search(father, i) != root {
			return false
		}
	}
	return true
}

func search(father []int, i int) int {
	son := i
	for father[i] != i {
		i = father[i]
		search(father, i)
	}

	// 路径压缩
	for son != i {
		tmp := father[son]
		father[son] = i
		son = tmp
	}

	return i
}
