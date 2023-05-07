package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("*****Подсчёт чисел в массиве*****")
	sizeArr := askSizeArr()
	emptyArr := createNewArr(sizeArr)
	fullArr := generateRandomNumInArr(emptyArr)
	printArr(fullArr)
	value := askNumForSearch(fullArr)
	amountNumInArrAfterValue := findAmountNumInArr(fullArr, value)
	printAmount(amountNumInArrAfterValue, value)
}

func askSizeArr() (size int) {
	fmt.Println("-----")
	fmt.Println("***Определяем размер массива***")
	fmt.Println("Введите размер массива:")
	fmt.Scan(&size)
	return
}

func createNewArr(size int) (arr []int) {
	fmt.Println("-----")
	fmt.Println("***Создаём пустой массив***")
	arr = make([]int, size)
	return
}

func generateRandomNumInArr(arr []int) []int {
	fmt.Println("-----")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10 * len(arr))
	}
	return arr
}

func printArr(arr []int) {
	fmt.Println("***Печать массива***")
	for _, el := range arr {
		fmt.Printf("%d\t", el)
	}
	fmt.Println()
	return
}

func askNumForSearch(arr []int) (value int) {
	fmt.Println("Введите число:")
	fmt.Scan(&value)
	return
}

func findAmountNumInArr(arr []int, value int) (amount int) {
	for idx, el := range arr {
		if el == value {
			amount = len(arr) - idx - 1
		}
	}
	return
}

func printAmount(amount, value int) {
	fmt.Println("-----")
	fmt.Printf("В массиве есть %d элементов после числа %d.\n", amount, value)
	return
}
