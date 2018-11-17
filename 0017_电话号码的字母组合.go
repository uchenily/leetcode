package leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	length := 1
	for _, digit := range digits {
		if digit == '9' || digit == '7' {
			length *= 4
		} else {
			length *= 3
		}
	}
	result := make([]string, length)
	count := length
	for _, digit := range digits {
		switch digit {
		case '2':
			data := []string{"a", "b", "c"}
			j := 0
			count /= 3
			for i := 0; i < length; i++ {
				result[i] += data[j]
				if (i+1)%count == 0 {
					j = (j + 1) % 3
				}
			}
			// fmt.Println(result)
		case '3':
			data := []string{"d", "e", "f"}
			j := 0
			count /= 3
			for i := 0; i < length; i++ {
				result[i] += data[j]
				if (i+1)%count == 0 {
					j = (j + 1) % 3
				}
			}
		case '4':
			data := []string{"g", "h", "i"}
			j := 0
			count /= 3
			for i := 0; i < length; i++ {
				result[i] += data[j]
				if (i+1)%count == 0 {
					j = (j + 1) % 3
				}
			}
		case '5':
			data := []string{"j", "k", "l"}
			j := 0
			count /= 3
			for i := 0; i < length; i++ {
				result[i] += data[j]
				if (i+1)%count == 0 {
					j = (j + 1) % 3
				}
			}
		case '6':
			data := []string{"m", "n", "o"}
			j := 0
			count /= 3
			for i := 0; i < length; i++ {
				result[i] += data[j]
				if (i+1)%count == 0 {
					j = (j + 1) % 3
				}
			}
		case '7':
			data := []string{"p", "q", "r", "s"}
			j := 0
			count /= 4
			for i := 0; i < length; i++ {
				result[i] += data[j]
				if (i+1)%count == 0 {
					j = (j + 1) % 4
				}
			}
		case '8':
			data := []string{"t", "u", "v"}
			j := 0
			count /= 3
			for i := 0; i < length; i++ {
				result[i] += data[j]
				if (i+1)%count == 0 {
					j = (j + 1) % 3
				}
			}
		case '9':
			data := []string{"w", "x", "y", "z"}
			j := 0
			count /= 4
			for i := 0; i < length; i++ {
				result[i] += data[j]
				if (i+1)%count == 0 {
					j = (j + 1) % 4
				}
			}
		}
	}
	return result
}
