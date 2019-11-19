import "math"

// 桶排序
// 构造的桶大小为 (B - A) / (N - 1)
// 则最大的 gap 必然是 当前桶里的最大值和下一个非空桶的最小值
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
