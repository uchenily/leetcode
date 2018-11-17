package leetcode

func twoSum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for index, value := range nums {
		m[value] = index
	}
	for index, value := range nums {
		if mValue, ok := m[target-value]; ok && index != mValue {
			return []int{index, mValue}
		}
	}
	return result
}
