package task_1624

import (
	"strings"
)

func maxLengthBetweenEqualCharacters(s string) int {

	letMap := make(map[string]int)
	for _, j := range s {
		letMap[string(j)]++
	}

	maxLen := -1

	for i, j := range letMap {
		minIndex := len(s)
		maxIndex := 0
		if j > 1 {
			minIndex = strings.Index(s, i)
			maxIndex = strings.LastIndex(s, i)
			if maxLen < maxIndex-minIndex {
				maxLen = maxIndex - minIndex
			}
		}
	}

	return maxLen
}
