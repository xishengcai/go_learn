package longestPalindrome

func palindromeCenter(s string) string {
	l := len(s)
	begin := 0
	maxLen := 0
	for i := 0; i < l; i++ {
		left := expend(s, i, i)
		if maxLen < 2*(i-left-1)+1 {
			maxLen = 2*(i-left-1) + 1
			begin = left + 1
		}

		if i+1 >= l || s[i] != s[i+1] {
			continue
		}
		left = expend(s, i, i+1)
		if maxLen < 2*(i-left-1)+2 {
			maxLen = 2*(i-left-1) + 2
			begin = left + 1
		}

	}
	return s[begin : begin+maxLen]
}

func expend(s string, left, right int) (halfLen int) {
	for left > -1 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	return left
}
