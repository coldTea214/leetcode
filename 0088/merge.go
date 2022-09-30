func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j int
	for i, j = m-1, n-1; i >= 0 && j >= 0; {
		k := i + j + 1
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	if j >= 0 {
		copy(nums1, nums2[:j+1])
	}
}
