package main

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	fmt.Println("*****27.8 Практическая работа*****\n")
	mapStudent := make(map[string]*Student)
	aboutStudents(mapStudent)
	printMap(mapStudent)
	fmt.Println("\nКонец программы")
}

func aboutStudents(mapStudent map[string]*Student) {
	for {
		nameStudent, ageStudent, gradeStudent := askNameAgeGradeStudent()
		if nameStudent == "" {
			break
		}
		personStruct := &Student{
			name:  nameStudent,
			age:   ageStudent,
			grade: gradeStudent,
		}
		mapStudent[personToString(personStruct)] = personStruct
	}
}

func askNameAgeGradeStudent() (string, int, int) {
	var nameStudent string
	var ageStudent, gradeStudent int
	fmt.Println("\tВведите имя студента:")
	fmt.Scan(&nameStudent)
	fmt.Println("\tВведите возраст студента:")
	fmt.Scan(&ageStudent)
	fmt.Println("\tВведите курс студента: ")
	fmt.Scan(&gradeStudent)
	fmt.Println()
	return nameStudent, ageStudent, gradeStudent
}

func personToString(p *Student) string {
	return fmt.Sprint(p.name, " ", p.age, p.grade)
}

func printMap(mapStudent map[string]*Student) {
	for key, _ := range mapStudent {
		fmt.Println(key)
	}
}
