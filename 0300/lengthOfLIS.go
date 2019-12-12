import "sort"

// O(n^2)
func lengthOfLIS(nums []int) int {
	// dp[i] 代表为第 i 个数字为结尾的最长上升子序列的长度
	dp, res := make([]int, len(nums)+1), 0
	for i := 1; i <= len(nums); i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// O(nlogn)
func lengthOfLIS(nums []int) int {
	// dp 代表当前最大上升序列长度为 len(dp)
	// dp[i] 代表当前长度下，最小的数字
	// eg: 输入 10 2 5 3，则 dp: [10]->[2]->[2,5]->[2,3]
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}
