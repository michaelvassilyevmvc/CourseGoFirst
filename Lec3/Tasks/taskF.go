package main

import "fmt"

func mainF() {

	var (
		first  string
		second string
		third  string
		fourth string
		name   string
	)

	fmt.Scan(&first, &second, &third, &fourth, &name)
	fmt.Printf("%s - %s\n", fourth, name)
	fmt.Printf("%s - %s\n", third, name)
	fmt.Printf("%s - %s\n", second, name)
	fmt.Printf("%s - %s\n", first, name)
}
