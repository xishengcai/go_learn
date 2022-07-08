package binary

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func(candidates []int, target int, rs []int, temp int)

	dfs = func(candidates []int, target int, rs []int, temp int) {
		// 解
		if temp == target {
			res = append(res, rs)
			return
		}

		// 无解
		if temp > target || len(candidates) == 0 {
			return
		}

		thisRs := append(rs, candidates[0])
		dfs(candidates[1:], target, thisRs, temp+candidates[0])
	}

	dfs(candidates, target, []int{}, 0)
	return res
}

func sum(x []int) int {
	var res int
	for _, k := range x {
		res += k
	}
	return res
}
