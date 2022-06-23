package binary

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(r, l *TreeNode) bool {

	if r == nil && l == nil {
		return true
	}

	if r == nil || l == nil {
		return false
	}

	return l.Val == r.Val && compare(r.Left, l.Right) && compare(r.Right, l.Left)
}
