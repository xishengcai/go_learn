package binary

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root.Val)

	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}

	return result
}

func inorderTraversal2(root *TreeNode) []int {
	var result []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inorder(root.Left)
		}
		result = append(result, root.Val)

		if root.Right != nil {
			inorder(root.Right)
		}
	}
	inorder(root)
	return result
}

func inorderTraversal3(root *TreeNode) {
	array := make([]*TreeNode, 0)
	for root != nil || len(array) > 0 {
		for root != nil {
			array = append(array, root)
			root = root.Left
		}
		if len(array) > 0 {
			root = array[len(array)-1]
			array = array[:len(array)-1]
			print(root.Val, " ")
			root = root.Right
		}
	}
	println("]")
}

// 线索二叉树
func inorderTraversal4(root *TreeNode) {
	var x *TreeNode
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
				print(root.Val, " ") // 右子节点指向root，打印
				root = root.Right    // 说明左子树已经遍历完成
			}
		} else {
			println(root.Val) // 左子节点为空，打印
			root = root.Right
		}
		// 结束条件,因为只会给左树添加右节点
	}
}
