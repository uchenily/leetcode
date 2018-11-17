package leetcode

// 递归解法, 没超时
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	left := findPeakElement(nums[:len(nums)/2])
	right := findPeakElement(nums[len(nums)/2:])
	result := left
	if nums[right+len(nums)/2] > nums[left] {
		result = right + len(nums)/2
	}
	return result
}

// 还有二分解法
