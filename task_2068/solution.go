package task_2068

import "math"

func checkAlmostEquivalent(word1 string, word2 string) bool {
	var letters1 = make(map[string]int)
	var letters2 = make(map[string]int)

	for _, let := range word1 {
		letters1[string(let)]++
	}

	for _, let := range word2 {
		letters2[string(let)]++
	}

	for let, _ := range letters1 {
		if (math.Abs(float64(letters1[let] - letters2[let]))) > 3 {
			return false
		}
	}

	for let, _ := range letters2 {
		if (math.Abs(float64(letters2[let] - letters1[let]))) > 3 {
			return false
		}
	}

	return true
}
