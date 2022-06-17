package treeSum

import (
	"sort"
)

//
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L := i + 1
		R := n - 1

		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				array := []int{nums[i], nums[L], nums[R]}
				result = append(result, array)

				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R--
			} else if nums[i]+nums[L]+nums[R] < 0 {
				if nums[L]+nums[i] > 0 {
					break
				}
				L++
			}
		}
	}
	println("{")
	for _, k := range result {
		print("[", k[0], " ", k[1], " ", k[2], "],")
	}
	println("}")
	return result
}

/*
[[-82,-11,93],[-82,13,69],[-82,17,65],[-82,21,61],[-82,26,56],
[-82,33,49],[-82,34,48],[-82,36,46],
[-70,-14,84],[-70,-6,76],
[-70,1,69],[-70,13,57],[-70,15,55],[-70,21,49],[-70,34,36],
[-66,-11,77],[-66,-3,69],[-66,1,65],[-66,10,56],[-66,17,49],

[-49,-6,55],[-49,-3,52],
[-49,1,48],[-49,2,47],[-49,13,36],
[-49,15,34],[-49,21,28],
[-43,-14,57],[-43,-6,49],[-43,-3,46],
[-43,10,33],[-43,12,31],[-43,15,28],[-43,17,26],[-29,-14,43]
,[-29,1,28],[-29,12,17],[-14,-3,17],[-14,1,13],[-14,2,12],
[-11,-6,17],[-11,1,10],[-3,1,2]]
*/
