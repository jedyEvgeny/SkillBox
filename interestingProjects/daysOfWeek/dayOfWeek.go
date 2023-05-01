package main

import (
	"fmt"
	"time"
)

func main() {
	//Для выбора локации:
	// loc, err := time.LoadLocation("Europe/Moscow")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	tm := time.Now() /*.In(loc)*/

	daysOfWeek := [7]string{
		"Воскресенье",
		"Понедельник",
		"Вторник",
		"Среда",
		"Четверг",
		"Пятница",
		"Суббота",
	}

	months := [13]string{
		"",
		"января",
		"февраля",
		"марта",
		"апреля",
		"мая",
		"июня",
		"июля",
		"августа",
		"сентября",
		"октября",
		"ноября",
		"декабря",
	}

	fmt.Printf("Сегодня %s, %d %s %d г., %02d:%02d:%02d\n",
		daysOfWeek[tm.Weekday()],
		tm.Day(),
		months[tm.Month()],
		tm.Year(),
		tm.Hour(),
		tm.Minute(),
		tm.Second())
}
