package task_709

import "strings"

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func toLowerCaseNative(s string) string {
	b := []byte(s)
	result := ""

	for _, letter := range b {
		if int(letter) < 91 && int(letter) > 64 {
			result += string(letter + 32)
		} else {
			result += string(letter)
		}
	}

	return result
}
