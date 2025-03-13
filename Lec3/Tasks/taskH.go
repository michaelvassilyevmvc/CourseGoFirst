package main

import (
	"fmt"
	"math"
)

func mainH() {

	var (
		a int
		b int
	)

	fmt.Scan(&a, &b)

	p := int(math.Pow(float64(a+b), 2))

	fmt.Print(p)
}
