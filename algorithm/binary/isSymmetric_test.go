package binary

import "testing"

func TestDuiChen(t *testing.T) {
	root := TreeNode{3,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
	}
	t.Log(isSymmetric(&root))
}
