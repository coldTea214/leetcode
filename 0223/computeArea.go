func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	res := area(A, B, C, D) + area(E, F, G, H)
	I, J, K, L := max(A, E), max(B, F), min(C, G), min(D, H)
	if I >= K || J >= L {
		return res
	}
	return res - area(I, J, K, L)
}

func area(A, B, C, D int) int {
	return (C - A) * (D - B)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
