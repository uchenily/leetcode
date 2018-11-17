package leetcode

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var i int
	for i = 0; i*i <= x; i++ {
	} // 这里从0开始遍历有点蠢, 最好是采用二分法
	return i - 1
}
