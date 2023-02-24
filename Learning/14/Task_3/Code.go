package main

import (
	"fmt"
)

func main() {
	fmt.Println("Задание 3. Именованные возвращаемые значения")
	fmt.Println("-----------------")
	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)
	fmt.Println("-----------------")
	fmt.Println(multiplication(num))
	fmt.Println(summation(num))
}

func multiplication(x int) (y int) {
	y = x * x
	return
}

func summation(x int) (y int) {
	y = x + x
	return
}
