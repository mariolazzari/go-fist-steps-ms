package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// init vars
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)

	// constants
	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)
	fmt.Println("Constants:", StatusOK, StatusConnectionReset, StatusOtherError)

	updateName(&firstName, "Mario")
	fmt.Println("New firstName:", firstName)

	// return more vars
	p, a := getSquareData(4)
	fmt.Printf("Perimeter = %f, Area = %f", p, a)
	// command line args
	fmt.Println("\n\nTotal command line args:", len(os.Args))
	for idx, arg := range os.Args {
		fmt.Printf("Index %v - %q\n", idx, arg)
	}
}

// function
func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

// update function argument via pointer
func updateName(name *string, newNane string) {
	*name = newNane
}

// return more vals
func getSquareData(l float32) (float32, float32) {
	p := 4 * l
	a := l * l

	return p, a
}
