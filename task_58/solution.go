package task_58

import "strings"

func lengthOfLastWord(s string) int {
	str := strings.Fields(s)
	return len(str[len(str)-1])
}
