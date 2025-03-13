package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println("Second line")
	// fmt.Print("First")
	// fmt.Print("Second")
	// fmt.Println("Third")
	// fmt.Printf("Hello, my name is %s\nMy age is %d\n","Bob",42)

	var height int = 183
	fmt.Println("My height is: ", height)

	var uid = 12345
	fmt.Println("My uid:", uid)

	var firstVar, secondVar = 20, 30
	fmt.Printf("First var: %d\nSecond var: %d\n", firstVar, secondVar)

	var (
		personName string = "Bob"
		personAge  int    = 42
		personUid  int
	)

	fmt.Printf("Name: %s\nAge: %d\nUid: %d\n", personName, personAge, personUid)

	count := 10
	fmt.Println("Count:", count)
	count++
	fmt.Println("Count:", count)

	width, length := 20.5, 30.5
	fmt.Printf("Min dimensional of rectangle is : %.3f\n", math.Min(width, length))
}
