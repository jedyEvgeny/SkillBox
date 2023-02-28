package main

import "fmt"

func main() {
	fmt.Println("Задание 1. Подсчёт чётных и нечётных чисел в массиве")
	fmt.Println("&&&&&&&&&&&&&&")
	var num [10]int
	for i, _ := range num {
		fmt.Println("Введите число")
		fmt.Scan(&num[i])
	}
	var countEven, countOdd int8
	for _, n := range num {
		if n%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}
	fmt.Printf("Количество чётных чисел: %d; количество нечётных чисел: %d.", countEven, countOdd)
}
