package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("some.txt")
	if err := os.Chmod("some.txt", 0444); err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer.WriteString("Текст")
	writer.Flush()
	// fileName := "Task_3.txt"
	// file, err := os.Create(fileName)
	// if err != nil {
	// 	fmt.Println("Не удалось создать файл")
	// 	panic(err)
	// 	return
	// }
	// defer file.Close()

	// if err = os.Chmod(fileName, 0444); err != nil {
	// 	panic(err)
	// 	return
	// }

	// var buf bytes.Buffer
	// buf.WriteString("Текст")
	// if _, err = file.Write(buf.Bytes()); err != nil {
	// 	fmt.Println("Не удалось выполнить запись в файл")
	// 	panic(err)
	// 	return
	// }
}
