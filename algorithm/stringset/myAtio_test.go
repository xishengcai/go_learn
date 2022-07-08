package stringset

import "testing"

func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("  123"))
	t.Log(myAtoi(" +123dd1"))
	t.Log(myAtoi(" -123 xx"))
	t.Log(myAtoi("words and 987"))

	t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))     // 49
	t.Log(maxArea([]int{1, 8, 100, 2, 100, 4, 8, 3, 7})) //200
}
