// 替换的规律是：
// 1. 从后往前，找到最长的降序排列
// 2. 降序转升序
// 3. 将序列前元素，与序列中第1个大于它的交换

//              2.           3.
// 1 2 3 4 5 -> 1 2 3 4 5 -> 1 2 3 5 4
// 1 2 5 4 3 -> 1 2 3 4 5 -> 1 3 2 4 5
// 1 3 2 5 4 -> 1 3 2 4 5 -> 1 3 4 2 5

func nextPermutation(a []int) {
	left := len(a) - 2
	for 0 <= left && a[left] >= a[left+1] {
		left--
	}

	reverse(a, left+1)

	// eg: 5 4 3 2 1
	if left == -1 {
		return
	}

	right := search(a, left+1, a[left])
	a[left], a[right] = a[right], a[left]
}

func reverse(a []int, l int) {
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

func search(a []int, l, target int) int {
	r := len(a) - 1
	l--
	for l+1 < r {
		mid := (l + r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
