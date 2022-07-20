package _6permute

import "sort"

func permute(nums []int) [][]int {

	res := [][]int{}

	sort.Ints(nums)

	var dfs func(nums, temp []int)

	dfs = func(nums, temp []int) {
		if len(nums) == 0 {
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] {
				continue
			}
			temp = append(temp, nums[i])
			var x []int
			for j := 0; j < len(nums); j++ {
				if j == i {
					continue
				}
				x = append(x, nums[j])
			}
			dfs(x, temp)
			var y []int
			for j := 0; j < len(temp)-1; j++ {
				y = append(y, temp[j])
			}

			temp = y

		}
	}
	t := []int{}
	dfs(nums, t)
	return res
}
