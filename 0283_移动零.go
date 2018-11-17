package leetcode

func moveZeroes(nums []int) {
	length := len(nums)
	pos := 0
	count0 := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		} else {
			count0++
		}
	}
	for i := count0; i > 0; i-- {
		nums[length-i] = 0
	}
}
