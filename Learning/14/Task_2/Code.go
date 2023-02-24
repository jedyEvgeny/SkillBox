package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Функция, возвращающая несколько значений")
	fmt.Println("---------------------")
	fmt.Println("Введите три числа:")
	var first, second, third int
	fmt.Scan(&first, &second, &third)
	randomFirstX, randomFirstY := randomNum(first)
	randomSecondX, randomSecondY := randomNum(second)
	randomThirdX, randomThirdY := randomNum(third)
	fmt.Println("---------------------")
	fmt.Println("Сгенерированные координаты первой точки:", randomFirstX, randomFirstY)
	fmt.Println("Сгенерированные координаты второй точки:", randomSecondX, randomSecondY)
	fmt.Println("Сгенерированные координаты третьей точки:", randomThirdX, randomThirdY)
	fmt.Println("---------------------")
	x1, y1 := coordinateTransform(randomFirstX, randomFirstY)
	x2, y2 := coordinateTransform(randomSecondX, randomSecondY)
	x3, y3 := coordinateTransform(randomThirdX, randomThirdY)

	fmt.Println("Преобразованные координаты первой точки:", x1, y1)
	fmt.Println("Преобразованные координаты второй точки:", x2, y2)
	fmt.Println("Преобразованные координаты третьей точки:", x3, y3)

}

func randomNum(a int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(a), rand.Intn(a)
}

func coordinateTransform(a, b int) (int, int) {
	a = 2*a + 10
	b = -3*b - 5
	return a, b
}
