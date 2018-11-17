package leetcode

func sortColors(nums []int) {
	left := -1
	right := len(nums)
	cursor := 0
	for cursor < right {
		if nums[cursor] == 1 {
			cursor++
		} else if nums[cursor] == 0 {
			nums[cursor], nums[left+1] = nums[left+1], nums[cursor]
			left++
			cursor++
		} else if nums[cursor] == 2 {
			nums[cursor], nums[right-1] = nums[right-1], nums[cursor]
			right--
		}
	}
}
