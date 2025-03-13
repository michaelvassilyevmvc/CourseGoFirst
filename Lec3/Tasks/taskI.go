package main

import "fmt"

func mainI() {
	var (
		a int
		b int
		c int
	)

	fmt.Scan(&a, &b, &c)
	fmt.Printf("%d%d%d", c, b, a)
}
