package leetcode

func searchRange(nums []int, target int) []int {
	index := findIndex(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}
	left, right := index, index
	for left >= 0 && nums[left] == target {
		left--
	}
	for right < len(nums) && nums[right] == target {
		right++
	}
	return []int{left + 1, right - 1}
}

// 二分查找寻找target的位置(其中一个, 可能不唯一), 没找到返回-1
func findIndex(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
