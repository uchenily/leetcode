package leetcode

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	result := 0
	for x > 0 {
		bit := x % 10
		result = result*10 + bit
		x /= 10
	}

	if result > 1<<31-1 || result <= -1<<31 {
		return 0
	}

	return result * sign
}
