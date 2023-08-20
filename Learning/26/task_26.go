package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("-----Программа cat----")
	err := checkAmountArgs()
	if err != nil {
		fmt.Println(err)
		// здесь можно прописать, как обрабатывать ошибку
	}
	fmt.Println("Программа cat закончила выполнение работы")
}

func checkAmountArgs() error {
	fmt.Println("---Определяем количество аргументов---")
	amountArgs := len(os.Args)
	if amountArgs > 1 && amountArgs < 5 {
		cutResult, err := oneArgs()
		if err != nil {
			fmt.Println(err)
			// здесь можно прописать, как обрабатывать ошибку
		}
		if amountArgs == 4 {
			threeArgs(cutResult)
		}
	} else {
		return fmt.Errorf("Пожалуйста, укажите три текстовых файла!")
	}
	return nil
}

func oneArgs() ([]byte, error) {
	fmt.Println("---Читаем содержимое файлов---")
	var catData []byte
	for i := 1; i < len(os.Args); i++ {
		file, err := os.ReadFile(os.Args[i])
		if err != nil {
			return catData, fmt.Errorf("Не удалось прочитать %d-й файл:\n", (i), err)
		}
		fmt.Println(string(file))
		catData = append(append(catData, "\n"...), file...)
	}
	return catData, nil
}

func threeArgs(data []byte) {
	err := ioutil.WriteFile(os.Args[3], []byte(data), 0777)
	if err != nil {
		fmt.Println("Не удалось записать в третий файл:", err)
		return
	}
}
