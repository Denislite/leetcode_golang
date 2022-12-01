package task_961

func repeatedNTimes(nums []int) int {
	count := len(nums) / 2
	for _, num := range nums {
		temp := 0
		for _, find := range nums {
			if num == find {
				temp++
			}
		}
		if temp == count {
			return num
		}
	}
	return 0
}
