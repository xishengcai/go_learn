package longestPalindrome

func palindromeDynamicPlan(s string) string {

	l := len(s)
	if l < 2 {
		return s
	}
	maxLen, begin := 1, 0
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = true
	}

	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			if s[j] == s[i] && (dp[j-1][i+1] || j-1 < i+1) {
				dp[j][i] = true
				if maxLen < j-i+1 {
					begin, maxLen = i, j-i+1
				}
			} else {
				dp[j][i] = false
			}
		}
	}

	return s[begin : begin+maxLen]
}
