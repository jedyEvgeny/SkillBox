package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var answer string
	timeTemplate := "01-02-2006 15:04:05"
	count := 1
	var writer bytes.Buffer

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
		writer.WriteString(fmt.Sprintf("%d %s %s\n", count, moment, answer))
		count++
	}
	if err = ioutil.WriteFile("Log.txt", writer.Bytes(), 0666); err != nil {
		panic(err)
		return
	}
	fmt.Println("Программа завершила работу")
}
