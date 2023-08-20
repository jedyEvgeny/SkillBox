package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("-----Программа cat----")
	checkAmountArgs()
	fmt.Println("Программа cat закончила выполнение работы")
}

func checkAmountArgs() {
	fmt.Println("---Определяем количество аргументов---")
	amountArgs := len(os.Args)
	if amountArgs > 1 && amountArgs < 5 {
		cutResult := oneArgs()
		if amountArgs == 4 {
			threeArgs(cutResult)
		}
	} else {
		fmt.Println("Пожалуйста, укажите три текстовых файла!")
		os.Exit(1)
	}
}

func oneArgs() []uint8 {
	fmt.Println("---Читаем содержимое файлов---")
	var catData []byte
	for i := 1; i < len(os.Args); i++ {
		file, err := os.ReadFile(os.Args[i])
		if err != nil {
			fmt.Printf("Не удалось прочитать %d-й файл:\n", (i), err)
			os.Exit(1)
		}
		fmt.Println(string(file))
		catData = append(append(catData, "\n"...), file...)
	}
	return catData
}

func threeArgs(data []uint8) {
	err := ioutil.WriteFile(os.Args[3], []uint8(data), 0777)
	if err != nil {
		fmt.Println("Не удалось записать в третий файл:", err)
		return
	}
}
