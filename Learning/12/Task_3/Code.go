package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// const timeTemplate = "2006-01-02 15:04:05"

func main() {
	fileName := "testfile.txt"
	fmt.Println("Создаем файл в:", time.Now())

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	fmt.Println("Файл создан в:", time.Now())
	fmt.Println("Записываем строку в файл в:", time.Now())
	line := fmt.Sprintf("%v %s\n", time.Now(), "Первая строка")

	if _, err = file.WriteString(line); err != nil {
		panic(err)
		return
	}
	fmt.Println("Строка записана в:", time.Now())
	fmt.Println("Изменяем права доступа на режим 'чтение' в:", time.Now())

	if err = os.Chmod(fileName, 0444); err != nil {
		panic(err)
		return
	}
	fmt.Println("Права доступа измены в:", time.Now())

	fmt.Println("Открываем файл на чтение в:", time.Now())
	if file, err = os.Open(fileName); err != nil {
		panic(err)
		return
	}
	defer file.Close()
	fmt.Println("Открыли файл для чтения. Пробуем прочитать в:", time.Now())

	info, err := file.Stat()
	if err != nil {
		panic(err)
		return
	}
	buf := make([]byte, info.Size())
	if _, err = file.Read(buf); err != nil {
		panic(err)
		return
	}
	fmt.Println("Файл прочитал в:", time.Now())
	fmt.Println("Вот что водержит файл:")
	fmt.Println("---------------------")
	fmt.Printf("%s", buf)
	fmt.Println("---------------------")

	fmt.Println("А можем ли мы что-нибудь записать в файл? Для начала откроем файл с модом на полный доступ, в том числе на запись")

	fmt.Println("Открываем файл на запись в:", time.Now())

	if file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777); err != nil {
		fmt.Println("Не удалось открыть файл с полным доступом в:", time.Now())
		panic(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	line = fmt.Sprintf("%v %s\n", time.Now(), "Вторая строка")

	if _, err = writer.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
		panic(err)
		return
	}

	writer.Flush()
	if err = writer.Flush(); err != nil {
		fmt.Println("Не удалось записать строку в файл")
		panic(err)
		return
	}

	fmt.Println("Новая строка записана в файл")
}
