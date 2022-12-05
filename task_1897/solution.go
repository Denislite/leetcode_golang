package task_1897

import "math"

func makeEqual(words []string) bool {
	var res = make(map[string]int)

	for _, str := range words {
		for _, let := range str {
			if string(let) != "" {
				res[string(let)]++
			}
		}
	}

	for k, _ := range res {
		if math.Mod(float64(res[k]), float64(len(words))) != 0 {
			return false
		}
	}

	return true
}
