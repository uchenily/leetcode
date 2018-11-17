package leetcode

func isAnagram(s string, t string) bool {
	cacheS := []byte(s)
	cacheT := []byte(t)
	if len(cacheS) != len(cacheT) {
		return false
	}

	sliceS := make([]int, 26)
	sliceT := make([]int, 26)
	for _, value := range cacheS {
		sliceS[value-'a']++
	}
	for _, value := range cacheT {
		sliceT[value-'a']++
	}

	// 比较
	for _, value := range cacheS {
		if sliceS[value-'a'] != sliceT[value-'a'] {
			return false
		}
	}
	return true
}
