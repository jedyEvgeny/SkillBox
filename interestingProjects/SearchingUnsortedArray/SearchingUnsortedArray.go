//Код поиска в несортированном массиве
//Сперва создаём массив и заполняем его случайными значениями
//Затем ищем в нём число перебором

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("*****Поиск по несортированному массиву*****")
	lenArr := tellLenArr()
	arr := createArray(lenArr)
	printArr(arr)
	findIdx(arr)
}

func tellLenArr() (outLenArr int) {
	fmt.Println("----------------------")
	fmt.Println("Введите длину массива:")
	fmt.Scan(&outLenArr)
	return
}

func createArray(inLenArr int) (outArr []int) {
	outArr = make([]int, inLenArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < inLenArr; i++ {
		outArr[i] = rand.Intn(100001)
	}
	return
}

func printArr(inArr []int) {
	for idx, el := range inArr {
		if idx+1 == len(inArr) {
			fmt.Printf("%d\n", el)
		} else {
			fmt.Printf("%d; ", el)
		}
	}
}

func findIdx(inArr []int) (outIdx int) {
	outIdx = -1
	fmt.Println("----------------------")
	fmt.Println("Введите число для поиска в массиве")
	ans := 0
	fmt.Scan(&ans)
	for i := 0; i < len(inArr); i++ {
		if ans == inArr[i] {
			outIdx = i
			break
		}
	}
	fmt.Printf("Индекс запрошенного числа %d равен %d.", ans, outIdx)
	return
}
