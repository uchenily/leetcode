package leetcode

import "sort"

// ByteSlice []byte
type ByteSlice []byte

// Len
func (a ByteSlice) Len() int {
	return len(a)
}

// Swap
func (a ByteSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less
func (a ByteSlice) Less(i, j int) bool {
	return a[i] < a[j]
}

func groupAnagrams(strs []string) [][]string {
	// 将排序后的字符串作为key, 原字符串作为value(中的一个)
	m := make(map[string][]string)
	for i := range strs {
		temp := ByteSlice(strs[i])
		sort.Sort(temp)
		key := string(temp)
		m[key] = append(m[key], strs[i])
	}
	// fmt.Println(m)
	result := make([][]string, len(m))
	i := 0
	for _, v := range m {
		result[i] = v
		i++
	}
	return result
}
