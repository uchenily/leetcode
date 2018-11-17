package leetcode

import "bytes"

func reverseString(s string) string {
	var buf bytes.Buffer
	bs := []byte(s)
	for i := len(bs) - 1; i >= 0; i-- {
		buf.WriteByte(bs[i])
	}
	return buf.String()
}
