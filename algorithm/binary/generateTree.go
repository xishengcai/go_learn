package binary

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return helperG(1, n)
}

func helperG(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	if start == end {
		return []*TreeNode{&TreeNode{start, nil, nil}}
	}
	var allTree []*TreeNode
	for i := start; i <= end; i++ {
		leftSubTree := helperG(start, i-1)
		rightSubTree := helperG(i+1, end)
		for _, l := range leftSubTree {
			for _, r := range rightSubTree {
				root := &TreeNode{i, nil, nil}
				root.Left, root.Right = l, r
				allTree = append(allTree, root)
			}
		}
	}
	return allTree
}
