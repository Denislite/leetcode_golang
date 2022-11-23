package task_120

import (
	"math"
)

//func main() {
//	a := [][]int{
//		{-1},
//		{2, 3},
//		{1, -1, -3},
//	}
//
//	fmt.Println(minimumTotal(a))
//}

func minimumTotal(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		for dot := range triangle[row] {
			triangle[row][dot] += int(math.Min(float64(triangle[row+1][dot]),
				float64(triangle[row+1][dot+1])))
		}
	}
	return triangle[0][0]
}
