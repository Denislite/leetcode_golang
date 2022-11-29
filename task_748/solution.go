package task_748

import (
	"regexp"
	"strings"
)

const pattern = "[a-z]"

func shortestCompletingWord(licensePlate string, words []string) string {
	re, _ := regexp.Compile(pattern)
	allFinds := re.FindAll([]byte(strings.ToLower(licensePlate)), -1)

	letterMap := make(map[string]int)

	for _, letter := range allFinds {
		letterMap[string(letter)]++
	}

	started := false
	result := ""

	for _, word := range words {
		good := true
		for letter, letterCount := range letterMap {
			if strings.Count(word, letter) < letterCount {
				good = false
			}
		}
		if !started && good {
			started = true
			result = word
		}
		if len(word) < len(result) && good {
			result = word
		}
	}

	return result
}
