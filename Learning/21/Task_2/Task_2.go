package main

import (
	"fmt"
)

func main() {
	fmt.Println(wrapFunc(20, 30, func(l, m int) int { return l + m }))
	fmt.Println(wrapFunc(20, 30, func(l, m int) int { return l * m }))
	fmt.Println(wrapFunc(20, 30, func(l, m int) int { return l - m }))
}

func wrapFunc(inL, inM int, A func(int, int) int) int {
	var result int
	defer func() {
		result = A(inL, inM)
		fmt.Println("Результат через defer равен:", result)
	}()
	return result
}
