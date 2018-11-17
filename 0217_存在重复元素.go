package leetcode

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for index, value := range nums {
		m[value] = index
	}
	if len(m) < len(nums) {
		return true
	}
	return false
}
