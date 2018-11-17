package leetcode

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	result := []int{}

	pos := 0
	for i := 0; i < len(nums1); i++ {
		for j := pos; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				pos = j + 1
				result = append(result, nums1[i])
				break
			}
		}
	}
	return result
}
