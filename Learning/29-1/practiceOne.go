package main

import (
	"fmt"
	"strconv"
	"sync"
)

const stop = "стоп"

var wg sync.WaitGroup

func main() {
	for {
		fmt.Print("Введите число или команду стоп и нажмите Enter: ")
		var ans string
		fmt.Scan(&ans)
		if ans == stop {
			fmt.Println("Программа завершена")
			return
		}
		num, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("введено не число и не стоп: ", err)
			continue
		}
		wg.Add(1)
		go square(num)
		wg.Wait()
	}
}

func square(numIn int) {
	defer wg.Done()
	sq := numIn * numIn
	fmt.Println("Квадрат: ", sq)
	wg.Add(1)
	go multiplication(sq)
}

func multiplication(sqIn int) {
	defer wg.Done()
	sqIn *= 2
	fmt.Println("Произведение: ", sqIn)
	fmt.Println("-----------")
}
