package leetcode

// 这种方法可以但是会超时
// func climbStairs(n int) int {
//     if n == 1 {
//         return 1
//     }
//     if n == 2 {
//         return 2
//     }
//     return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	// result[i] 表示到达第i级阶梯的方法数
	result := make([]int, n+1)
	result[1] = 1
	result[2] = 2
	for i := 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}
