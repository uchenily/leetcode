package leetcode

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	result := true
	length := len(s)
	s = strings.ToLower(s)
	fmt.Println(s)
	i, j := 0, length-1
	for i < j {
		if !isLetterOrNumber(s[i]) {
			i++
			continue
		}
		if !isLetterOrNumber(s[j]) {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
		} else {
			result = false
			break
		}
	}
	return result
}

func isLetterOrNumber(b byte) bool {
	if b <= '9' && b >= '0' || b <= 'z' && b >= 'a' || b <= 'Z' && b >= 'A' {
		return true
	}
	return false
}
