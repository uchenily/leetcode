package leetcode

func firstUniqChar(s string) int {
	slice := make([]int, 26)
	cache := []byte(s)
	for _, value := range cache {
		slice[value-'a']++
	}
	// fmt.Println(slice)
	for index, value := range cache {
		if slice[value-'a'] == 1 {
			return index
		}
	}
	return -1
}
