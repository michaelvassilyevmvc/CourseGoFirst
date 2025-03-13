package main

import "fmt"

func mainD() {
	var (
		name     string
		lastName string
		age      int
		category string = "Студент BPS"
	)

	fmt.Scan(&name, &lastName, &age)
	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . %s", name, lastName, age, category)

}
