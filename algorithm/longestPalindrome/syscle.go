package longestPalindrome

// 暴力循环法
func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	begin := 0
	maxLen := 0
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			// 判断回文法
			if judge(s, i, j) && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}

// judge
func judge(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
