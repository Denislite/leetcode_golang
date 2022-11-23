package task_136

func singleNumber(nums []int) int {
	for _, num := range nums {
		var count int
		for _, secondNum := range nums {
			if num == secondNum {
				count++
			}
		}
		if count == 1 {
			return num
		}
	}
	return -1
}
