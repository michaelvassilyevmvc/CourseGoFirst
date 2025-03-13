package main

import (
	"fmt"
)

func mainJ() {
	var (
		a float32
		b float32
	)

	fmt.Scan(&a, &b)
	sum := (int)(a + b)
	if sum%2 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}
}
