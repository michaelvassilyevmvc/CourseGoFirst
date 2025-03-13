package main

import "fmt"

func mainG() {

	var (
		a int
		b int
	)

	fmt.Scan(&a, &b)
	p := 2 * (a + b)
	s := a * b

	fmt.Printf("Периметр: %d\n", p)
	fmt.Printf("Площадь: %d\n", s)
}
