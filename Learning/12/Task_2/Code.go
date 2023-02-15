package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("Log.txt")
	if err != nil {
		fmt.Println("Файл отсутствует")
		panic(err)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
		return
	}
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	if info.Size() == 0 {
		fmt.Println("Файл пуст")
		return
	}
	fmt.Println(string(buf))
}
