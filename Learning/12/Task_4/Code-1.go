package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	file, err := os.Create("Log.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
	var writer bytes.Buffer

	for {
		fmt.Println("Введите запрос:")
		var answer string
		fmt.Scan(&answer)
		if answer == "exit" {
			break
		}
		timeNow := time.Now()
		writer.WriteString(timeNow.Format("01-02-2006 15:04:05 "))
		writer.WriteString(fmt.Sprintf("%s\n", answer))
	}
	if err = ioutil.WriteFile("Log.txt", writer.Bytes(), 0666); err != nil {
		panic(err)
		return
	}
	fmt.Println("Программа завершила работу")
}
