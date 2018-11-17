package leetcode

func titleToNumber(s string) int {
	str := []byte(s)
	result := 0
	for _, v := range str {
		result = result*26 + int(v) - 'A' + 1
	}
	return result
}
