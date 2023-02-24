package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите число")
	var answer int
	fmt.Scan(&answer)
	fmt.Println(evenUneven(answer))
}

func evenUneven(a int) (b bool) {
	if a%2 == 0 {
		b = true
	} else {
		b = false
	}
	return
}
