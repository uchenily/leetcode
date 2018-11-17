package leetcode

import "fmt"

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	// sums[i]表示到第i家时最多能偷到的金额
	sums := make([]int, length)
	sums[0] = nums[0]
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] && i >= 2 {
			sums[i] = sums[i-2] + nums[i]
		} else if nums[i] > nums[i-1] { // 第二家, 且比第一家多
			sums[i] = nums[i]
		} else if i >= 2 {
			sums[i] = maxOne(sums[i-1], nums[i]+sums[i-2])
		} else { // 第二家, 且比第一家少
			sums[i] = sums[i-1]
		}
	}
	fmt.Println(sums)
	return maxElem(sums)
}

func maxElem(s []int) int {
	max := s[0]
	for i := 0; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}
	return max
}

func maxOne(a, b int) int {
	if a > b {
		return a
	}
	return b
}
