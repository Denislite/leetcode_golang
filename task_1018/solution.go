package task_1018

func prefixesDivBy5(nums []int) []bool {
	var result []bool

	temp := 0
	for _, i := range nums {
		temp = (temp*2 + i) % 5
		if temp == 0 {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
