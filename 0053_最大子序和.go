package leetcode

func maxSubArray(nums []int) int {
	length := len(nums)
	sums := make([]int, length)
	for i := 0; i < length; i++ {
		sums[i] = -1000
	}
	// sums[n]表示以第n个元素结尾且和最大的连续子数组和
	sums[0] = nums[0]
	for i := 1; i < length; i++ {
		if sums[i-1] > 0 {
			sums[i] = sums[i-1] + nums[i]
		} else {
			sums[i] = nums[i]
		}
	}
	return max(sums)
}

func max(s []int) int {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
