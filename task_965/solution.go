package task_965

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	leaves := *checkLeaves(root, &[]int{})

	check := leaves[0]
	for _, num := range leaves {
		if num != check {
			return false
		}
	}

	return true
}

func checkLeaves(root *TreeNode, leaves *[]int) *[]int {
	if root == nil {
		return leaves
	}

	checkLeaves(root.Left, leaves)
	*leaves = append(*leaves, root.Val)
	checkLeaves(root.Right, leaves)

	return leaves
}
