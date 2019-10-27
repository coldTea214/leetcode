func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	collection := make([]byte, n)
	for i := 0; i < n; i++ {
		collection[i] = byte(i) + '1'
	}
	permutationNum := 1
	for i := 2; i < n; i++ {
		permutationNum *= i
	}
	k--

	res := make([]byte, n)
	for i := 0; i < n-1; i++ {
		idx := k / permutationNum
		res[i] = collection[idx]
		collection = append(collection[:idx], collection[idx+1:]...)
		k %= permutationNum
		permutationNum /= (n - i - 1)
	}
	res[n-1] = collection[0]

	return string(res)
}
