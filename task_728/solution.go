package task_728

import "math"

func selfDividingNumbers(left int, right int) []int {
	var result []int
	for i := left; i <= right; i++ {
		check := true
		temp := i
		for temp > 0 {
			if math.Mod(float64(i), math.Mod(float64(temp), 10)) != 0 {
				check = false
				break
			}
			temp /= 10
		}
		if check {
			result = append(result, i)
		}
	}
	return result
}
