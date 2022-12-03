package task_2053

func kthDistinct(arr []string, k int) string {
	var strs = make(map[string]int)

	for _, str := range arr {
		strs[str]++
	}

	var dist []string
	for _, str := range arr {
		if strs[str] == 1 {
			dist = append(dist, str)
		}
	}

	if len(dist) < k {
		return ""
	}

	return dist[k-1]
}
