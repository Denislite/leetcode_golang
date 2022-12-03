package task_2057

import "math"

func smallestEqual(nums []int) int {
	for i, num := range nums {
		if int(math.Mod(float64(i), 10)) == num {
			return i
		}
	}
	return -1
}
