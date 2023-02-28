package main

import "fmt"

func main() {
	fmt.Println("Задание 2. Функция, реверсирующая массив")
	fmt.Println("________________________")
	var (
		arr        [10]int
		arrReverse [10]int
	)
	count := len(arr) - 1
	for i, _ := range arr {
		fmt.Println("Введите элемент массива:")
		fmt.Scan(&arr[i])
		arrReverse[count] = arr[i]
		count--
	}
	for _, num := range arrReverse {
		fmt.Printf("%d ", num)
	}
}
