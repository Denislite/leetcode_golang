package task_26

func removeDuplicates(nums []int) int {
	find := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == find {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
		if nums[i] != find {
			find = nums[i]
		}
	}
	return len(nums)
}
