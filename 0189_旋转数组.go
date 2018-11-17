package leetcode

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)
	k %= length
	for i := 0; i < (length-k)/2; i++ {
		nums[i], nums[length-1-k-i] = nums[length-1-k-i], nums[i]
	}
	fmt.Println(length-k, nums)
	for i := 0; i < k/2; i++ {
		nums[length-k+i], nums[length-1-i] = nums[length-1-i], nums[length-k+i]
	}
	fmt.Println(nums)
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
	}
}
