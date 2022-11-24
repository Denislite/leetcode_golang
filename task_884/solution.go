package task_884

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	str := strings.Split(s1, " ")
	str = append(str, strings.Split(s2, " ")...)

	var checked = make(map[string]int)

	for _, s := range str {
		checked[s]++
	}

	var result []string

	for st, count := range checked {
		if count == 1 {
			result = append(result, st)
		}
	}

	return result
}
