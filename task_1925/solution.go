package task_1925

func countTriples(n int) int {
	var nums = make(map[int]int)
	for i := 1; i <= n; i++ {
		nums[i*i] = i
	}

	result := 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if _, check := nums[i*i+j*j]; check {
				result += 2
			}
		}
	}
	return result
}
