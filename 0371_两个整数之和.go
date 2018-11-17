package leetcode

func getSum(a int, b int) int {
	if a&b == 0 {
		return a | b
	}
	return getSum(a^b, (a&b)<<1)
}
