package main

import (
	"fmt"
)

func main() {
	var (
		X1 int
		Y1 int
		X2 int
		Y2 int
	)

	fmt.Scan(&X1, &Y1, &X2, &Y2)

	if (X2-X1 == 2 && Y2-Y1 == 1) ||
		(X2-X1 == 2 && Y2-Y1 == -1) ||
		(X1-X2 == 2 && Y2-Y1 == 1) ||
		(X1-X2 == 2 && Y2-Y1 == -1) ||

		(Y2-Y1 == 2 && X2-X1 == 1) ||
		(Y2-Y1 == 2 && X2-X1 == -1) ||
		(Y1-Y2 == 2 && X2-X1 == 1) ||
		(Y1-Y2 == 2 && X2-X1 == -1) {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
