package task_125

import (
	"regexp"
	"strings"
)

const reg = "[a-z0-9]"

func isPalindrome(s string) bool {

	re, _ := regexp.Compile(reg)
	allFinds := re.FindAll([]byte(strings.ToLower(s)), -1)

	var result, revert string

	for i, _ := range allFinds {
		result += string(allFinds[i])
		revert += string(allFinds[len(allFinds)-i-1])
	}
	if result == revert {
		return true
	}
	return false
}
