package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	var result, temp [][]int
	temp = permute(nums[1:])
	result = make([][]int, len(temp)*len(nums))
	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(nums); j++ {
			// 这里有点绕, 其实就是将第一个数依次插入到前一个结果的不同位置
			result[i*len(nums)+j] = append(result[i*len(nums)+j], temp[i][:j]...)
			result[i*len(nums)+j] = append(result[i*len(nums)+j], nums[0])
			result[i*len(nums)+j] = append(result[i*len(nums)+j], temp[i][j:]...)
		}
	}
	return result
}

/**
1

[1] 2
2 [1]

[1] 2 3
2 [1] 3
2 3 [1]
[1] 3 2
3 [1] 2
3 2 [1]

...
*/
