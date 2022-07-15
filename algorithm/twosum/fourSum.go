package twosum

import "sort"

func fourSum(nums []int, target int) [][]int {

	l := len(nums)

	result := [][]int{}

	sort.Ints(nums)
	for i := 0; i <= l-4; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[l-3]+nums[l-2]+nums[l-1] < target {
			continue
		}
		for j := i + 1; j <= l-3; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := nums[i] + nums[j]
			L := j + 1
			R := l - 1
			for L < R {
				if sum+nums[L]+nums[R] == target {
					array := []int{nums[i], nums[j], nums[L], nums[R]}
					result = append(result, array)
					for L < R && nums[L] == nums[L+1] {
						L++
					}
					for L < R && nums[R] == nums[R-1] {
						R--
					}
					L++
					R--
				} else if sum+nums[L]+nums[R] > target {
					R--
				} else if sum+nums[L]+nums[R] < target {
					L++
				}
			}
		}
	}
	return result
}
