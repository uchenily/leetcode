package leetcode

import "fmt"

func canJump(nums []int) bool {
	length := len(nums)
	// dp[i] 表示第i个位置是否可以到达
	dp := make([]bool, length)
	if length == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	dp[0] = true
	for i := 0; i < length; i++ {
		for j := 0; j < nums[i]; j++ {
			if dp[i] == true && i+j+1 < length {
				dp[i+j+1] = true
			}
		}
	}
	fmt.Println(dp)
	return dp[length-1]
}
