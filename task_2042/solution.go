package task_2042

import (
	"fmt"
	"regexp"
	"strconv"
)

const pattern = "[0-9]*"

func areNumbersAscending(s string) bool {
	re, _ := regexp.Compile(pattern)
	allFinds := re.FindAll([]byte(s), -1)

	var nums []int
	for _, num := range allFinds {
		if string(num) != "" {
			str, _ := strconv.Atoi(string(num))
			nums = append(nums, str)
		}
	}

	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}
