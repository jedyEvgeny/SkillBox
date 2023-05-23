package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-----25.8 Практическая работа-----")
	str, substr := createFlags()
	if checkFlagsExist(str, substr) != true {
		return
	}
	searchSubstringInString(str, substr)
}

func createFlags() (str, substr *string) {
	fmt.Println("\n---Назначаем и считываем флаги из CMD---")
	// Назначаем и считываем флаги как аргументы из командной строки
	str = flag.String("str", "", "Кириллическая строка, в которой будет выполнен поиск")
	substr = flag.String("substr", "", "Кириллическая подстрока для поиска")
	flag.Parse()
	return
}

func checkFlagsExist(str, substr *string) bool {
	fmt.Println("\n---Проверяем наличие двух флагов---")
	if *str == "" || *substr == "" {
		fmt.Println("Ошибка: обязательные аргументы -str и -substr не предоставлены.")
		return false
	}
	return true
}

func searchSubstringInString(str, substr *string) {
	fmt.Println("\n---Ищем подстроку в строке---")
	index := strings.Index(*str, *substr)
	if index != -1 {
		//вместо грустного вывода true or false, вывожу расширенное сообщение
		fmt.Printf("Ура! :) подстрока '%s' найдена в строке '%s' начиная с позиции %d.\n", *substr, *str, index)
	} else {
		fmt.Printf("Увы, (: подстрока '%s' не найдена в строке '%s'.\n", *substr, *str)
	}
}
