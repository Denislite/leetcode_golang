package task_1636

import "sort"

func frequencySort(nums []int) []int {
	intMap := make(map[int]int)
	for _, j := range nums {
		intMap[j]++
	}
	sort.Slice(nums, func(a, b int) bool {
		if intMap[nums[a]] == intMap[nums[b]] {
			return nums[a] > nums[b]
		}
		return intMap[nums[a]] < intMap[nums[b]]
	})

	return nums
}
