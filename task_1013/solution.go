package task_1013

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	if sum%3 != 0 {
		return false
	}

	count := sum / 3

	tempCount := 0
	temp := 0
	for _, num := range arr {
		temp += num
		if temp == count {
			temp = 0
			tempCount++
			if tempCount == 3 {
				return true
			}
		}
	}

	return false
}
