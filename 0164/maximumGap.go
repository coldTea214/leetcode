import "math"

// 桶排序
// 构造的桶大小为 (Max - Min) / (N - 1), 设 Max=0,Min=10,N=6, 则桶为 [0,2) [2,4) [4,6) [6,8) [8,10) [10,12)
// 则maxGap必然是 当前桶里的最大值和下一个非空桶的最小值
// 反证: 如上例, 假如maxGap位于一个桶内, 则maxGap<2(桶左闭右开), 6个数之间的gap就算都为maxGap, 也是maxGap*5<Max-Min
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	max, min := math.MinInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	bucketLen := int(math.Ceil(float64(max-min) / float64(len(nums)-1)))
	if bucketLen == 0 {
		return 0
	}

	bucketNum := (max-min)/bucketLen + 1
	minInBucket := make([]int, bucketNum)
	maxInBucket := make([]int, bucketNum)
	for i := 0; i < bucketNum; i++ {
		minInBucket[i] = math.MaxInt32
		maxInBucket[i] = math.MinInt32
	}
	for i := 0; i < len(nums); i++ {
		idx := (nums[i] - min) / bucketLen
		if nums[i] < minInBucket[idx] {
			minInBucket[idx] = nums[i]
		}
		if nums[i] > maxInBucket[idx] {
			maxInBucket[idx] = nums[i]
		}
	}

	// 引入maxInBucketBefore因为最大gap可能跨多个桶
	res, maxInBucketBefore := 0, maxInBucket[0]
	for i := 1; i < bucketNum; i++ {
		if minInBucket[i] == math.MaxInt32 {
			continue
		}

		if minInBucket[i]-maxInBucketBefore > res {
			res = minInBucket[i] - maxInBucketBefore
		}
		maxInBucketBefore = maxInBucket[i]
	}
	return res
}
