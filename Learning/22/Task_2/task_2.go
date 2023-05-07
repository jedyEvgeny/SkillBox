package main

import (
	"fmt"
)

const (
	sizeArr = 12
)

func main() {
	fmt.Println("*****Нахождение первого вхождения числа в упорядоченном массиве*****")
	emptyArr := createArr()
	fullArrRandomNum := fillArrRandomNum(emptyArr)
	printArr(fullArrRandomNum)
	needNum := findNum()
	indexNeedNum := findIdx(fullArrRandomNum, needNum)
	printIdx(indexNeedNum, needNum)
}

func createArr() (outArr [sizeArr]int) {
	fmt.Println("----------")
	fmt.Println("***Создаём массив***")
	outArr = [sizeArr]int{}
	return
}

func fillArrRandomNum(arr [sizeArr]int) [sizeArr]int {
	fmt.Println("----------")
	fmt.Println("***Введите через пробел значения элементов массива. В конце нажмите Enter***")
	var ans int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&ans)
		arr[i] = ans
	}
	return arr
}

func printArr(arr [sizeArr]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\t", arr[i])
		if i == len(arr)-1 {
			fmt.Println() //Следующая печать с новой строки
		}
	}
	return
}

func findNum() (outNum int) {
	fmt.Println("----------")
	fmt.Println("Введите число для поиска:")
	fmt.Scan(&outNum)
	return
}

func findIdx(arr [sizeArr]int, num int) (idx int) {
	fmt.Println("***Поиск в массиве***")
	idx = -1
	max := len(arr) - 1
	min := 0
	for max >= min {
		mid := (max + min) / 2
		if num == arr[mid] {
			idx = mid
			break
		} else if num < arr[max] {
			max = mid - 1
			continue
		} else {
			min = mid + 1
		}
	}
	for i := 0; i < idx; i++ {
		if arr[i] == num {
			idx = i
		}

	}
	return
}

func printIdx(idx int, num int) {
	fmt.Printf("Число %d в массиве имеет индекс %d.", num, idx)
	return
}
