package binary

/*
Question Description:

Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

Note:
A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {

	var x *TreeNode
	var y *TreeNode
	var pre *TreeNode
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inorder(root.Left)
		}
		if pre == nil {
			pre = root
		} else {
			if pre.Val > root.Val && x == nil {
				x = pre
				y = root
			} else if pre.Val > root.Val {
				y = root
				return
			}
			pre = root
		}
		if root.Right != nil {
			inorder(root.Right)
		}
	}
	inorder(root)

	x.Val, y.Val = y.Val, x.Val
}

func recoverTree3(root *TreeNode) {
	var x, l, r, pre *TreeNode
	for root != nil {
		// 遍历左子树，找到最右边的节点，添加后继节点
		if root.Left != nil {
			x = root.Left
			for x.Right != nil && x.Right != root {
				// 防止无限循环
				x = x.Right
			}
			// 左子树的，结束条件：
			if x.Right == nil {
				x.Right = root   // 添加后驱，root
				root = root.Left // 继续下一个左子树
			} else {
				if pre != nil && root.Val < pre.Val {
					r = root
					if l == nil {
						l = pre
					}
				}

				x.Right = nil
				pre = root
				root = root.Right // 说明左子树已经遍历完成
			}
		} else {
			if pre != nil && root.Val < pre.Val {
				r = root
				if l == nil {
					l = pre
				}
			}
			pre = root
			root = root.Right

		}
	}
	if l != nil && r != nil {
		l.Val, r.Val = r.Val, l.Val
	}
}
