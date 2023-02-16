package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	fileName := "file1.txt"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	if err = os.Chmod(fileName, 0400); err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file) //убедился, что проверка уровня доступа осуществялется при открытии файла
	if err != nil {
		fmt.Println(err)
		return
	}

	writer.WriteString("Текст")
	writer.Flush()

	if file, err = os.Open(fileName); err != nil {
		panic(err)
		return
	}
	defer file.Close()

	var s string
	var writerTwo bytes.Buffer
	fmt.Println("Введите слово:")
	fmt.Scan(&s)
	writerTwo.WriteString(s)

	if _, err = file.Write(writerTwo.Bytes()); err != nil {
		fmt.Println("Не удалось записать в файл, т.к. не позволяет уровень доступа")
		panic(err)
		return
	}
}
