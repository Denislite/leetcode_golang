package task_9

import (
	"strconv"
)

func isPalindrome(x int) bool {
	return strconv.Itoa(x) == reverse(strconv.Itoa(x))
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
