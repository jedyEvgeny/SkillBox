package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var answer string
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
		timeNow := time.Now()
		file.WriteString(timeNow.Format("01-02-2006 15:04:05 "))
		file.WriteString(fmt.Sprintf("%s\n", answer))
	}
	fmt.Println("Программа завершила работу")
}
