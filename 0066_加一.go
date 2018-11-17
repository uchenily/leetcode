package leetcode

func plusOne(digits []int) []int {

	// 如果所有位都是9,才会多一位
	result := true
	for _, value := range digits {
		if value != 9 {
			result = false
			break
		}
	}
	if result {
		temp := make([]int, len(digits)+1)
		for i := 0; i < len(digits); i++ {
			temp[i+1] = digits[i]
		}
		digits = temp
	}

	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 10 {
			digits[i-1]++
			digits[i] = 0
		}
	}
	return digits
}
