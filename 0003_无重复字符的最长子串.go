package leetcode

func lengthOfLongestSubstring(s string) int {
	index := 0 // 记录当前最长子串的起始下标
	strs := []byte(s)
	// result[i] 表示以第i个元素结尾的最长子串长度
	length := len(strs)
	if length <= 1 {
		return length
	}
	result := make([]int, len(strs))
	result[0] = 1
	for i := 1; i < len(strs); i++ {
		var j int
		for j = index; j < i; j++ {
			if strs[i] == strs[j] {
				result[i] = i - j
				index = j + 1
				break
			}
		}
		if j == i {
			result[i] = result[i-1] + 1
		}
	}
	// fmt.Println(result)

	var max int
	for _, v := range result {
		if max < v {
			max = v
		}
	}
	return max
}

// 另一种思路, 使用一个map保存字符最后出现的位置, start指示当前无重复字符串的开头位置
// 如果这个字符出现的位置在start前, 长度加一
// 出现在start后, 则需要重新定位start
