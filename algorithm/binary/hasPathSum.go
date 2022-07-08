package binary

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) [][]int {

	var res [][]int

	var dep func(root *TreeNode, targetSum int, arr []int)

	dep = func(root *TreeNode, targetSum int, arr []int) {
		if root == nil {
			return
		}

		targetSum -= root.Val
		arr = append(arr, root.Val)

		if targetSum == 0 && root.Left == nil && root.Right == nil {
			res = append(res, append([]int(nil), arr...))
			return
		}

		dep(root.Left, targetSum, arr)
		dep(root.Right, targetSum, arr)
		return

	}
	dep(root, targetSum, make([]int, 0))
	return res
}

func flatten(root *TreeNode) {
	x := &TreeNode{}
	temp := x
	var dep func(root *TreeNode)

	dep = func(root *TreeNode) {
		if root == nil {
			return
		}

		temp.Right = &TreeNode{
			root.Val,
			nil,
			nil,
		}
		temp = temp.Right

		dep(root.Left)
		dep(root.Right)
	}

	dep(root)

	root.Right = x.Right.Right
}
