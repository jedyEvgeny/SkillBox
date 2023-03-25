package main

import (
	"fmt"
)

const (
	rows = 3
	cols = 3
)

func main() {
	matrix := [rows][cols]int{
		{15, 26, 13},
		{16, 5, 167},
		{22, 64, 365},
	}
	fmt.Println(detMatrix(matrix))
}

func detMatrix(inMatrix [rows][cols]int) int {
	lDetTriangleOne := inMatrix[0][0] * inMatrix[1][1] * inMatrix[2][2]
	lDetTriangleTwo := inMatrix[0][1] * inMatrix[1][2] * inMatrix[2][0]
	lDetTriangleThree := inMatrix[1][0] * inMatrix[2][1] * inMatrix[0][2]
	lDet := lDetTriangleOne + lDetTriangleTwo + lDetTriangleThree

	rDetTriangleOne := inMatrix[0][2] * inMatrix[1][1] * inMatrix[2][0]
	rDetTriangleTwo := inMatrix[2][1] * inMatrix[1][2] * inMatrix[0][0]
	rDetTriangleThree := inMatrix[1][0] * inMatrix[0][1] * inMatrix[2][2]
	rDet := rDetTriangleOne + rDetTriangleTwo + rDetTriangleThree

	det := lDet - rDet

	return det
}
