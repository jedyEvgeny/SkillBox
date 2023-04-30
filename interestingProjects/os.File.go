package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "Defer.txt"
	file := createFile(fileName)
	defer closeFile(file)
	writeFile(file)
}

func createFile(inFileName string) *os.File {
	log.Println("Создание файла.")
	inFile, err := os.Create(inFileName)
	if err != nil {
		log.Fatalf("Ошибка при создании файла: %v.", err)
	}
	log.Printf("Файл %v создан.\n", inFileName)
	return inFile
}

func writeFile(inFile *os.File) {
	log.Println("Запись в файл")
	// fmt.Fprintln(inFile, "Записанные данные")
	fmt.Println("Введите текст для записи в файл:")
	var ans string
	fmt.Scan(&ans)
	writer := bufio.NewWriter(inFile)
	if _, err := writer.WriteString(ans); err != nil {
		log.Fatalf("Ошибка при записи в файл: %v.", err)
	}
	if err := writer.Flush(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func closeFile(file *os.File) {
	log.Println("Закрываем файл")
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v.\n", err)
		os.Exit(1)
	}
}
