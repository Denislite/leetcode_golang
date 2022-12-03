package task_2032

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var nums = make(map[int]int)

	nums1 = removeDuplicateValues(nums1)
	nums2 = removeDuplicateValues(nums2)

	nums1 = append(nums1, nums2...)

	nums3 = removeDuplicateValues(nums3)

	nums1 = append(nums1, nums3...)

	for _, num := range nums1 {
		nums[num]++
	}

	var result []int
	for num, _ := range nums {
		if nums[num] > 1 {
			result = append(result, num)
		}
	}

	return result
}

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
