package task_509

func fib(n int) int {
	if n == 0 {
		return 0
	}

	y := 1
	x := 0

	for i := 1; i < n; i++ {
		y = x + y
		x = y - x
	}

	return y
}
