package task_1539

func findKthPositive(arr []int, k int) int {
	mis, j, num := 0, 0, 1
	find := false

	for !find {
		if num != arr[j] {
			mis++
		} else {
			if j+1 < len(arr) {
				j++
			}
		}
		if mis == k {
			find = true
			return num
		}
		num++
	}
	return 0
}
