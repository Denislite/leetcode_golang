package task_938

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	leaves := *checkLeaves(root, &[]int{})

	result := 0
	for _, num := range leaves {
		if num >= low && num <= high {
			result += num
		}
	}

	return result
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
