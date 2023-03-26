package main

import (
	"fmt"
)

const (
	rowsOne = 3
	colsOne = 5
	rowsTwo = 5
	colsTwo = 4
)

func main() {
	matrixOne := [rowsOne][colsOne]int{
		{15, 26, 13, 22, -88},
		{16, 5, 167, 22, 168},
		{22, 64, 365, -329, -682},
	}
	matrixTwo := [rowsTwo][colsTwo]int{
		{112, 62, 14, 54},
		{12, -62, 24, 36},
		{115, 8, -14, 21},
		{1, -55, 29, 38},
		{6, -2, 17, 222},
	}
	fmt.Println(multMatrix(matrixOne, matrixTwo))
}

func multMatrix(inMatrixA [rowsOne][colsOne]int, inMatrixB [rowsTwo][colsTwo]int) (outMatrix [rowsOne][colsTwo]int) {
	for k := 0; k < len(inMatrixA); k++ { //По вертикали
		for i := 0; i < len(inMatrixB[0]); i++ { //По горизонтали
			for j := 0; j < len(inMatrixA[0]); j++ { //Сложение произведений
				outMatrix[k][i] += inMatrixA[k][j] * inMatrixB[j][i]
			}
		}
	}
	return
}
