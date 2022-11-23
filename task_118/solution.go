package task_118

func generate(numRows int) [][]int {
	var result [][]int

	for i := 1; i <= numRows; i++ {
		numb := make([]int, i)
		numb[0] = 1
		numb[i-1] = 1
		result = append(result, numb)
	}
	for i, row := range result {
		for j, num := range row {
			if num == 0 {
				result[i][j] = result[i-1][j] + result[i-1][j-1]
			}
		}
	}
	return result
}
