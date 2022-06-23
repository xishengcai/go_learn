package bucket

import "testing"

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	case1 := []int{1, 5, 9, 1, 5, 9}
	k := 2
	tt := 3
	t.Log("Except: ", false)
	t.Log(containsNearbyAlmostDuplicate3(case1, k, tt))
}
