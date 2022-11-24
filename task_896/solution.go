package task_896

func isMonotonic(nums []int) bool {
	checked := false
	incr := true

	for i := 1; i < len(nums); i++ {
		if !checked {
			if nums[i] < nums[i-1] {
				incr = false
			} else if nums[i] == nums[i-1] {
				continue
			}
			checked = true
		}
		if incr {
			if nums[i] < nums[i-1] {
				return false
			}
		} else {
			if nums[i] > nums[i-1] {
				return false
			}
		}
	}

	return true
}
