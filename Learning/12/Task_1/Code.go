package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var answer string
	timeTemplate := "01-02-2006 15:04:05"
	count := 1

	file, err := os.Create("Log.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	for {
		fmt.Println("Введите запрос:")
		fmt.Scan(&answer)
		if answer == "exit" {
			break
		}
		moment := time.Now().Format(timeTemplate)
		file.WriteString(fmt.Sprintf("%d %s %s\n", count, moment, answer))
		count++
	}
	fmt.Println("Программа завершила работу")
}
