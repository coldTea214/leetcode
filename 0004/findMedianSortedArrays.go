// 执行用时 : 16 ms , 在所有 golang 提交中击败了 88.44% 的用户
// 内存消耗 : 5.5 MB , 在所有 golang 提交中击败了 79.21% 的用户

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)&1 != 0 {
		return float64(findKth((m+n)/2+1, nums1, nums2))
	} else {
		tmp := findKth((m+n)/2, nums1, nums2) + findKth((m+n)/2+1, nums1, nums2)
		return float64(tmp) / 2
	}
}

// 找第k小的数
func findKth(k int, nums1 []int, nums2 []int) int {
	if len(nums1) > len(nums2) {
		return findKth(k, nums2, nums1)
	}
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	i, j := min(k/2, len(nums1))-1, min(k/2, len(nums2))-1
	if nums1[i] > nums2[j] {
		return findKth(k-j-1, nums1, nums2[j+1:])
	} else {
		// 关键是对于 nums1[i] == nums2[j] 的分析
		// 这种情况实际是无所谓的，干掉哪一半都行，就先把短的数组清空理论上更快
		return findKth(k-i-1, nums1[i+1:], nums2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
