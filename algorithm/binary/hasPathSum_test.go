package binary

import "testing"

func TestHasPathSum(t *testing.T) {
	root := &TreeNode{5,
		&TreeNode{4,
			&TreeNode{11,
				&TreeNode{7, nil, nil},
				&TreeNode{2, nil, nil}},
			nil},
		&TreeNode{8,
			&TreeNode{13, nil, nil},
			&TreeNode{4,
				&TreeNode{5, nil, nil},
				&TreeNode{1, nil, nil}}}}
	t.Log(hasPathSum(root, 22))

	// [1,2,5,3,4,null,6]
	root2 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{5, nil,
			&TreeNode{6, nil, nil}}}
	flatten(root2)

}
