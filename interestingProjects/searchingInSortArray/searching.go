// В коде мы создаём массив
// Заполняем его случайными числами
// Сортируем массив
// Вводим число для поиска индекса в массиве
// Ищем число и выводим в терминал индекс. Если числа в массиве нет, то индекс равен -1
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("*****Поиск в сортированном массиве*****")
	emptyArr := createArr()
	printArr(emptyArr)
	fullArrRandomNum := fillArrRandomNum(emptyArr)
	printArr(fullArrRandomNum)
	orderingArr := orderArr(fullArrRandomNum)
	printArr(orderingArr)
	needNum := findNum()
	indexNeedNum := findIdx(orderingArr, needNum)
	printIdx(indexNeedNum, needNum)
}

func createArr() (outArr []int) {
	fmt.Println("----------")
	fmt.Println("***Создаём массив***")
	fmt.Println("Введите размер массива")
	var sizeArr int
	fmt.Scan(&sizeArr)
	outArr = make([]int, sizeArr)
	return
}

func fillArrRandomNum(arr []int) []int {
	fmt.Println("----------")
	fmt.Println("***Заполняем массив случайными значениями***")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(20)
	}
	return arr
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\t", arr[i])
		if i == len(arr)-1 {
			fmt.Println() //Следующая печать с новой строки
		}
	}
	return
}

func orderArr(arr []int) []int {
	fmt.Println("***Сортируем массив по порядку***")
	sort.Ints(arr)
	return arr
}

func findNum() (outNum int) {
	fmt.Println("----------")
	fmt.Println("Введите число для поиска:")
	fmt.Scan(&outNum)
	return
}

func findIdx(arr []int, num int) (idx int) {
	fmt.Println("***Поиск в массиве***")
	idx = -1
	max := len(arr) - 1
	min := 0
	for max >= min {
		mid := (max + min) / 2
		if num == arr[mid] {
			idx = mid
			return
		} else if num < arr[max] {
			max = mid - 1
			continue
		} else {
			min = mid + 1
		}
	}
	return
}

func printIdx(idx int, num int) {
	fmt.Printf("Число %d в массиве имеет индекс %d.", num, idx)
	return
}
