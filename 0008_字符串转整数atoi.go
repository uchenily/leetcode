package leetcode

func myAtoi(str string) int {
	cache := []byte(str)
	pos := 0
	sign := 1
	result := 0

	// 跳出空白字符
	for pos < len(cache) && cache[pos] == ' ' {
		pos++
	}

	// 判断是否是'+-'符号
	if pos < len(cache) {
		if cache[pos] == '-' {
			sign = -1
			pos++
		} else if cache[pos] == '+' {
			pos++
		}
	}

	// 数字判断
	for ; pos < len(cache); pos++ {
		if cache[pos] <= '9' && cache[pos] >= '0' {
			result = result*10 + int(cache[pos]-'0')
			if result*sign > 1<<31-1 {
				return 1<<31 - 1
			} else if result*sign < -1<<31 {
				return -1 << 31
			}
		} else {
			break
		}

	}
	result *= sign
	return result
}
