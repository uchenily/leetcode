package leetcode

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// first, second表示"当前"最小值, 右边比最小值大的一个数
	first, second := 1<<31-1, 1<<31-1 // INT_MAX
	for _, v := range nums {
		if v <= first {
			first = v
		} else if v <= second {
			second = v
		} else {
			return true
		}
		// fmt.Println("first", first, "second", second)
	}

	return false
}
