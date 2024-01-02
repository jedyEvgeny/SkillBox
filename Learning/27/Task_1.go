package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

type Storage map[string]*Student

func newStudent(name string, age int, grade int) *Student {
	return &Student{name, age, grade}
}

func (s Storage) put(name string, student *Student) {
	s[name] = student
}

func (s Storage) get() {
	for _, value := range s {
		fmt.Println(value.name, value.age, value.grade)
	}
}

func main() {
	storage := make(Storage)
	getInfoAboutAllStudent(storage)
	fmt.Println("--------------------")
	fmt.Println("Имена всех студентов:")
	fmt.Println("--------------------")
	storage.get()
}

func getInfoAboutAllStudent(storage Storage) {
	for {
		input := genInfoAboutOneStudent()
		if input == "" {
			break
		}
		data, msg := getArrAboutStudent(input)
		if msg != "" {
			continue
		}
		name, age, grade, msg := separateInfoAboutStudent(data)
		if msg != "" {
			fmt.Println(msg)
			continue
		}
		student := newStudent(name, age, grade)
		storage.put(name, student)
	}
}

func genInfoAboutOneStudent() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите данные студента (имя возраст класс) или нажимте Enter для завершения: ")
	scanner.Scan()
	input := scanner.Text()
	return input
}

func getArrAboutStudent(input string) (data []string, msg string) {
	data = strings.Split(input, " ")
	if len(data) != 3 {
		msg = "Неверный формат ввода"
		fmt.Println(msg)
	}
	return
}

func separateInfoAboutStudent(info []string) (string, int, int, string) {
	var msg string
	name := info[0]
	age, err := strconv.Atoi(info[1])
	if err != nil {
		msg = "Ошибка в возрасте"
	}
	grade, errGrade := strconv.Atoi(info[2])
	if errGrade != nil {
		msg = "Ошибка в классе"
	}
	if err != nil && errGrade != nil {
		msg = "Ошибка в возрасте и классе"
	}
	return name, age, grade, msg
}
