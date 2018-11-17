package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	index := 0
	for index < len(nums) && nums[index] <= 0 { // 固定一个数nums[index]
		left := index + 1
		right := len(nums) - 1
		for left < right {
			if nums[index]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[index], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] { // 跳过重复元素
					left++
				}
				for left < right && nums[right] == nums[right+1] { // 跳过重复元素
					right--
				}
			} else if nums[index]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[index]+nums[left]+nums[right] < 0 {
				left++
			}
		}
		index++
		for index > 0 && index < len(nums) && nums[index] == nums[index-1] { // 跳过重复元素
			index++
		}
	}
	return result
}
