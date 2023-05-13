package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sizeArr = 10

func main() {
	arr := createArr()
	printArr(arr)
	// Сортировка вставками
	start := time.Now()
	insertionSort(arr)
	duration := time.Since(start)
	printArr(arr)
	fmt.Println("Cортировка вставками O(n^2) занимает:", duration)
}

func createArr() (arr []int) {
	arr = make([]int, sizeArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)
	}
	return
}

func printArr(inArr []int) {
	for _, el := range inArr {
		fmt.Printf("%v\t", el)
	}
	fmt.Println()
	fmt.Println("-------------")
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
