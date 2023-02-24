package main

import (
	"fmt"
)

func main() {
	fmt.Println("Функция, принимающая значения по ссылке")
	fmt.Println("---------------------")
	var a, b int
	fmt.Println("Введите два числа")
	fmt.Scan(&a, &b)
	fmt.Println("---------------------")
	fmt.Println("Вы ввели числа:", a, b)
	changeSite(&a, &b)
	fmt.Println("---------------------")
	fmt.Println("После смены значений получаем:")
	fmt.Println(a, b)
}

func changeSite(x, y *int) {
	*x, *y = *y, *x
}
