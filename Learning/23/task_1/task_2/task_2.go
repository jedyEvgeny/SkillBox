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

	bubbleSortDescending := func(arr []int) {
		if len(arr) == 0 {
			return
		}
		for i := 0; i < len(arr)-1; i++ {
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	bubbleSortDescending(arr)
	printArr(arr)

}

func createArr() (arr []int) {
	arr = make([]int, sizeArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr * 100)
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
