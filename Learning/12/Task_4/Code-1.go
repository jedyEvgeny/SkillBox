package main

import (
	"bytes"
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
	var writer bytes.Buffer

	for {
		fmt.Println("Введите запрос:")
		fmt.Scan(&answer)
		if answer == "exit" {
			break
		}
		timeNow := time.Now()
		writer.WriteString(timeNow.Format("01-02-2006 15:04:05 "))
		writer.WriteString(fmt.Sprintf("%s\n", answer))
	}
	if _, err = file.Write(writer.Bytes()); err != nil {
		panic(err)
		return
	}
	fmt.Println("Программа завершила работу")
}
