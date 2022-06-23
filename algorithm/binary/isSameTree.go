package binary

/*
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if p.Left != nil || q.Left != nil {
		if !isSameTree(p.Left, q.Left) {
			return false
		}
	}

	if p.Right != nil || q.Right != nil {
		if !isSameTree(p.Right, q.Right) {
			return false
		}
	}

	return true
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	// 代码需要优化
	//if (p == nil && q != nil) || (p != nil && q == nil) {
	//	return false
	//}

	// 走到这里说明肯定有一个不为nil
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	return isSameTree(p.Right, q.Right)
}
