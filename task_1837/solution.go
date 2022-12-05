package task_1837

func sumBase(n int, k int) int {
	result := 0

	for ; n != 0; n /= k {
		result += n % k
	}

	return result
}
