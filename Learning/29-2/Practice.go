package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal)
	numbers := make(chan int)

	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go getNaturalNum(numbers)
	go getSquare(numbers, ch)
	<-ch
	fmt.Println("\nПолучен сигнал ^C. Выхожу из программы")
}

func getNaturalNum(numbers chan<- int) {
	count := 1
	for {
		numbers <- count
		count++
	}
}

func getSquare(numbers <-chan int, ch <-chan os.Signal) {
	for {
		select {
		case num := <-numbers:
			square := num * num
			fmt.Println("Квадрат числа", num, "равен:", square)
		case <-ch:
			return
		}
	}
}
