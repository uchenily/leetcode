package leetcode

func removeDuplicates(nums []int) int {
	number := 0
	for _, value := range nums {
		if value != nums[number] {
			number++
			nums[number] = value
		}
	}
	number++
	return number
}
