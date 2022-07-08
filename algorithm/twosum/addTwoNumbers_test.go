package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{1, 2, 3, -1, 0, 2, 1}, 3))
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{Val: 4, Next: &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{3, nil}}}}
	t.Log(addTwoNumbers(l1, l2))

	l3 := &ListNode{9, &ListNode{Val: 9, Next: &ListNode{9, nil}}}
	l4 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	t.Log(addTwoNumbers(l3, l4))
}
