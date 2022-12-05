package task_1903

import "math"

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if math.Mod(float64(num[i]), 2) == 1 {
			return num[:i+1]
		}
	}
	return ""
}
