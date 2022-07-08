package twosum

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
链表
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail, head *ListNode
	var x int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = &ListNode{}
		}

		if l2 == nil {
			l2 = &ListNode{}
		}
		current := l1.Val + l2.Val + x
		next := &ListNode{
			Val: current % 10,
		}
		if head == nil {
			head = next
			tail = head
		} else {
			tail.Next = next
			tail = next
		}
		x = current / 10
		l1, l2 = l1.Next, l2.Next
	}
	if x != 0 && tail != nil {
		tail.Next = &ListNode{x, nil}
	}
	return head
}
