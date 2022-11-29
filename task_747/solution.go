package task_747

func dominantIndex(nums []int) int {
	max, index := maxNum(nums)
	nums[index] = 0
	for _, num := range nums {
		if num*2 > max {
			return -1
		}
	}
	return index
}

func maxNum(nums []int) (int, int) {
	result := 0
	index := 0
	for i, num := range nums {
		if num > result {
			result = num
			index = i
		}
	}
	return result, index
}
