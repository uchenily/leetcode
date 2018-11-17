package leetcode

func trailingZeroes(n int) int {
	// 因式分解, n! = x*10^k = x*2^m*5^k,(m 一定大于 k) 只需要看因式分解后5出现多少次!
	sum := 0
	for n > 0 {
		sum += n / 5
		n /= 5
	}
	return sum
}
