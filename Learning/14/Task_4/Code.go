package main

import (
	"fmt"
)

const ConstA = 10
const ConstB = 100
const ConstC = 1000

func main() {
	fmt.Println("Задание 4. Область видимости переменных")
	fmt.Println("--------------------")
	fmt.Println("Введите число:")
	var num int
	fmt.Scan(&num)
	sum1 := sumFirst(num)
	sum2 := sumSecond(sum1)
	sum3 := sumThird(sum2)
	fmt.Println("--------------------")
	fmt.Println("Результат работы программы:", sum1, sum2, sum3)
}

func sumFirst(a int) int {
	a = a + ConstA
	return a
}

func sumSecond(a int) int {
	a = a + ConstB
	return a
}

func sumThird(a int) int {
	a = a + ConstB
	return a
}
