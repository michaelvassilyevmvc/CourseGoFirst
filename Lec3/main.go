package main

import "fmt"

func main() {
	var (
		age  int
		name string
	)

	fmt.Scan(&age)
	fmt.Scan(&name)

	fmt.Printf("My name is: %s\nMy age is : %d\n", name, age)
}
