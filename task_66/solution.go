package task_66

func plusOne(digits []int) []int {
	last := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			last = 0
			break
		} else {
			digits[i] = 0
			last = 1
		}
	}
	if last != 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
