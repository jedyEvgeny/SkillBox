package main

import (
	"fmt"
	"io/ioutil"
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

	if info.Size() == 0 {
		fmt.Println("Файл пуст")
		return
	}

	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(resultBytes))
}

