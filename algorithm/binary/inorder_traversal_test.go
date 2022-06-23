package binary

import "testing"

func TestFront(t *testing.T) {
	root := TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, &TreeNode{3, nil, nil}, nil}
	t.Log(inorderTraversal(&root))
	inorderTraversal3(&root)
	inorderTraversal4(&root)

	//root2 := TreeNode{1, &TreeNode{4,nil,nil}, nil}
	//t.Log(isSameTree2(&root, &root2))
	//t.Log(isSameTree2(&root, &root))
}

// 验证是否是搜索二叉树
func TestIsValidBTS(t *testing.T) {
	// [5,4,6,nul,nul,3,7]
	root := TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{
		6,
		&TreeNode{3, nil, nil},
		&TreeNode{7, nil, nil},
	}}

	root2 := TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{
		4,
		&TreeNode{3, nil, nil},
		&TreeNode{6, nil, nil},
	}}
	t.Log(isValidBST(&root))
	t.Log(isValidBST(&root2))
}

func TestRecover(t *testing.T) {
	// [3,1,4,null,null,2]
	root := TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{
		4,
		&TreeNode{2, nil, nil},
		nil,
	}}
	//[1,3,null,null,2]
	// 3-2-1
	// 2-3-1
	// 2-1-3
	root2 := TreeNode{1, &TreeNode{3, nil, &TreeNode{2, nil, nil}}, nil}

	recoverTree(&root)
	recoverTree3(&root2)
	recoverTree3(&root)
}
