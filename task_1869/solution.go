package task_1869

import (
	"strings"
)

func checkZeroOnes(s string) bool {
	maxOne, maxZero := 0, 0
	ones := strings.Split(s, "0")
	zeros := strings.Split(s, "1")
	for _, one := range ones {
		if one != "" && len(one) > maxOne {
			maxOne = len(one)
		}
	}
	for _, zero := range zeros {
		if zero != "" && len(zero) > maxZero {
			maxZero = len(zero)
		}
	}
	return maxOne > maxZero
}
