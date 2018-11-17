package leetcode

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	// maxPowerOfThree := 1
	// for maxPowerOfThree < 1<<32 {
	// 	maxPowerOfThree *= 3
	// }
	maxPowerOfThree := 10460353203
	// fmt.Println(maxPowerOfThree)
	return maxPowerOfThree%n == 0
}
