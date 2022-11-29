package task_744

func nextGreatestLetter(letters []byte, target byte) byte {

	count := 121
	result := letters[0]

	for _, letter := range letters {
		if int(letter-target) < count && int(letter-target) > 0 {
			result = letter
			count = int(letter - target)
		}
	}

	return result
}
