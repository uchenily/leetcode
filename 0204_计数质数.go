package leetcode

// 优化版本
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	// notPrimes[i]表示第i个数不是质数
	// 初始假设所有数都是质数, 后续逐步剔除
	notPrimes := make([]bool, n)
	result := n - 2 // 质数个数1,...n-1, 先除去2, 剩下n-2个
	for i := 2; i < n; i++ {
		if !notPrimes[i] { // i是质数, 将2i,3i,...剔除
			for j := 2 * i; j < n; j += i {
				if !notPrimes[j] { // 未标记过
					result--
					notPrimes[j] = true
				}
			}
		}
	}
	return result
}
