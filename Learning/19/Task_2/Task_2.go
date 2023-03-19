package main

import (
	"fmt"
)

const sizeArr = 6

func main() {
	fmt.Println("Сортировка пузырьком")
	arr := [sizeArr]int{12, 64, 52, 329, 4867, 79}
	fmt.Println(bubbleArr(arr))
	fmt.Println("The end")
}

func bubbleArr(inputArr [sizeArr]int) [sizeArr]int {
	var countChanges int
	for countChanges != sizeArr {
		countChanges = sizeArr
		for i := 0; i < sizeArr-1; i++ {
			if inputArr[i] > inputArr[i+1] {
				inputArr[i], inputArr[i+1] = inputArr[i+1], inputArr[i]
				countChanges--
			}
		}
	}
	return inputArr
}
