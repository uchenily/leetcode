package leetcode

func strStr(haystack string, needle string) int {
	pos := -1
	if needle == "" || haystack == needle {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	sliceH := []byte(haystack)
	sliceN := []byte(needle)
	for i := range sliceH {
		pos = i
		for j, val := range sliceN {
			if pos < len(haystack) {
				if val != sliceH[pos] {
					break
				} else {
					pos++
				}
			} else {
				return -1
			}

			if j == len(sliceN)-1 {
				return pos - len(sliceN)
			}
		}
	}

	return -1
}
