package palindromic_substring

func Solution(s string) string {
	length := len(s)
	if length == 0 {
		return s
	}
	pd := make([][]bool, length)
	// for i := 0; i < length; i++ {
	// 	pd[i] = make([]bool, length)
	// }
	leftIndex, rightIndex := 0, 0
	for i := 0; i < length; i++ {
		pd[i] = make([]bool, length)
		start, end := i, i
		for {
			if start < 0 {
				break
			}
			if start == end {
				pd[start][end] = true
			} else if start+1 == end {
				pd[start][end] = s[start] == s[end]
			} else {
				pd[start][end] = s[start] == s[end] && pd[start+1][end-1]
			}

			if pd[start][end] && rightIndex-leftIndex < end-start {
				leftIndex = start
				rightIndex = end
			}

			start--
		}

	}
	return s[leftIndex : rightIndex+1]
}
