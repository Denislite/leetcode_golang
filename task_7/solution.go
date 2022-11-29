package task_7

import "strconv"

func reverse(x int) int {
	num := 1
	if x < 0 {
		x *= -1
		num = -1
	}

	result, _ := strconv.Atoi(reverseStr(strconv.Itoa(x)))
	return num * result
}

func reverseStr(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
