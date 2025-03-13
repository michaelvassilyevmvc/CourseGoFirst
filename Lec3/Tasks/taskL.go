package main

import (
	"fmt"
	"strings"
)

func mainL() {
	var (
		login string
		email string
	)

	fmt.Scan(&login)
	if strings.Contains(login, "@") || len(login) < 10 {
		fmt.Println("Некорректный логин")
	} else {
		fmt.Scan(&email)
		if strings.Contains(email, "@") && strings.Contains(email, ".") {
			fmt.Println("ОК")
		} else {
			fmt.Println("Некорректная почта")
		}
	}

}
