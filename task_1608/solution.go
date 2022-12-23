package task_1636

import "sort"

func specialArray(nums []int) int {
	result := -1

	sort.Slice(nums, func(a, b int) bool {
		return a < b
	})

	for i := 0; i <= len(nums); i++ {
		count := 0
		for _, num := range nums {
			if count > i {
				break
			} else if num >= i {
				count++
			}
		}
		if count == i {
			result = count
		}
	}

	return result
}
