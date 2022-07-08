package binary

func numTrees(n int) int {

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		// root = i, left * right(笛卡尔积)
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]

		}
	}
	return dp[n]
}
