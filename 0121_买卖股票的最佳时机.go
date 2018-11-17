package leetcode

func maxProfit(prices []int) int {
	length := len(prices)
	max := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if prices[j] > prices[i] && prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
