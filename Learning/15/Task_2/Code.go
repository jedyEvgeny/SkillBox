package main

import (
	"fmt"
)

func main() {
	fmt.Println("Задание 2. Функция, реверсирующая массив")
	fmt.Println("________________________")
	var arr [10]int
	for i, _ := range arr {
		fmt.Println("Введите элемент массива:")
		fmt.Scan(&arr[i])
	}
	arr = revers(arr)
	fmt.Println(arr)
}

func revers(array [10]int) [10]int {
	for i, el := range array {
		array[len(array)-i-1] = el
	}
	return array
}
