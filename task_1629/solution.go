package task_1629

func slowestKey(releaseTimes []int, keysPressed string) byte {
	index := 0
	maxRes := releaseTimes[0]
	maxLet := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		tempRes := releaseTimes[i] - releaseTimes[i-1]
		if tempRes > maxRes {
			index = i
			maxRes = releaseTimes[i] - releaseTimes[i-1]
			maxLet = keysPressed[i]
		} else if tempRes >= maxRes && keysPressed[i] > maxLet {
			index = i
			maxRes = releaseTimes[i] - releaseTimes[i-1]
			maxLet = keysPressed[i]
		}
	}

	return keysPressed[index]
}
