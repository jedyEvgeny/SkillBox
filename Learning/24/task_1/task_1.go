package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("*****Задание 1. Чётные и нечётные*****")
	arr := createArr()
	fmt.Println("Исходный слайс:")
	printArr(arr)
	fmt.Println("------------")
	evenArr, oddArr := separateArr(arr)
	fmt.Println("Слайс чётных чисел:")
	printArr(evenArr)
	fmt.Println("------------")
	fmt.Println("Слайс нечётных чисел:")
	printArr(oddArr)
	ans := ""
	for ans != "exit" {
		fmt.Println("Для выхода из программы наберите exit")
		fmt.Scan(&ans)
	}
}

func createArr() []int {
	fmt.Println("Создаём слайс из случайных чисел")
	fmt.Println("------------")
	fmt.Println("Введите размер слайса:")
	sizeArr := 0
	fmt.Scan(&sizeArr)
	outArr := make([]int, sizeArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sizeArr; i++ {
		outArr[i] = rand.Intn(sizeArr * 100)
	}
	return outArr
}

func printArr(inArr []int) {
	for _, el := range inArr {
		fmt.Printf("%v\t", el)
	}
	fmt.Println()
	fmt.Printf("Количество элементов слайса: %v, ёмкость слайса: %v.\n", len(inArr), cap(inArr))
}

func separateArr(arr []int) ([]int, []int) {
	var evenArr []int
	var oddArr []int
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			evenArr = append(evenArr, arr[i])
		} else {
			oddArr = append(oddArr, arr[i])
		}
	}
	return evenArr, oddArr
}
