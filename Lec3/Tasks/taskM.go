package main

import (
	"fmt"
)

func mainM() {
	var (
		one   string
		two   string
		three string
	)

	fmt.Scan(&one, &two, &three)
	if ((one == "раз" || one == "один") && two == "два" && three == "три") || (one == "1" && two == "2" && three == "3") {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}
}
