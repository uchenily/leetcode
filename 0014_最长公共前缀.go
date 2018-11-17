package leetcode

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length < 1 {
		return ""
	} else if length == 1 {
		return strs[0]
	}
	pos := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < length; j++ {
			if len(strs[j]) > i && strs[0][i] == strs[j][i] {
				if j == length-1 {
					pos++
				}
			} else {
				return strs[0][:pos]
			}
		}
	}
	return strs[0][:pos]
}
