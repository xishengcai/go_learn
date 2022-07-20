package _1nextPermutation

import "testing"

func TestNextPermutation(t *testing.T) {
	nextPermutation([]int{1, 2, 5, 4, 3, 0})
	nextPermutation([]int{1, 2})
	nextPermutation([]int{3, 2, 1})
}
