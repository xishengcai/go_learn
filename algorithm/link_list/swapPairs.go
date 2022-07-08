package link_list

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre, next := head, head.Next
	result := head.Next
	for pre != nil && next != nil {
		pre, next = next, pre
		next.Next = pre.Next
		pre.Next = next

		pre = next.Next
		if next.Next != nil && next.Next.Next != nil {
			next.Next = next.Next.Next
			next = pre.Next
		} else {
			next = nil
		}

	}
	return result
}

// 递归
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead

}
