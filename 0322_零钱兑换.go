package leetcode

func coinChange(coins []int, amount int) int {
	// dp[i]表示凑成i块钱的最少硬币数
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	// 动态规划
	// 状态转移方程
	// dp[i] = min(dp[i-coins[j]])+1
	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				if dp[i] == -1 || dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	return dp[amount]
}
