package back

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func(candidates []int, target int, rs []int)

	dfs = func(candidates []int, target int, rs []int) {
		// 解
		if target == 0 {
			res = append(res, append([]int(nil), rs...))
			return
		}
		if target < 0 {
			return
		}

		for i, k := range candidates {
			x := target - k
			if x < 0 {
				continue
			}
			rs := append(rs, k)
			dfs(candidates[i:], x, rs)
		}

	}

	dfs(candidates, target, []int{})
	return res
}
