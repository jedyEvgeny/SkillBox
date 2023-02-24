package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Введите два числа")
	fmt.Scan(&a, &b)
	sumRange(a, b)
}

func sumRange(x, y int) {
	if x > y {
		x, y = y, x
	}
	var count int
	for x != y {
		if x%2 == 0 {
			count += x
		}
		x++
	}
	fmt.Println(count)

}
