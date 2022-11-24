package task_897

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	leaves := *checkLeaves(root, &[]int{})

	sort.Slice(leaves, func(i, j int) bool {
		return leaves[i] > leaves[j]
	})

	result := &TreeNode{
		Val:   leaves[0],
		Right: nil,
		Left:  nil,
	}

	for i := 1; i < len(leaves); i++ {
		leave := &TreeNode{
			Val:   leaves[i],
			Right: result,
			Left:  nil,
		}
		result = leave
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
