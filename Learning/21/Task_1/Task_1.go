package main

import (
	"fmt"
)

func main() {
	var x int16
	var y uint8
	var z float32
	fmt.Println("Введите три числа")
	fmt.Scan(&x, &y, &z)
	fmt.Println(sum(x, y, z))
}

func sum(x int16, y uint8, z float32) float32 {
	return 2*float32(x) + float32(y*y) - 3/z
}
