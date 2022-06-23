package binary

import "testing"

func TestIsSameTree(t *testing.T) {
	root := TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}
	root2 := TreeNode{1, &TreeNode{4, nil, nil}, nil}
	t.Log(isSameTree(&root, &root2))
	t.Log(isSameTree2(&root, &root))
}
