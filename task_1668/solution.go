package task_1668

import "strings"

func maxRepeating(sequence string, word string) int {
	count := 0
	index := 0
	k := 1
	for index > -1 {
		index = strings.Index(sequence, strings.Repeat(word, k))
		if index > -1 {
			count++
		}
		k++
	}
	return count
}
