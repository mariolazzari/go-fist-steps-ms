package main

import "fmt"

func main() {
	val := 0
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &val)
	if val < 0 {
		panic("Negative numbers not allowed")
	}
	fmt.Println("You entered:", val)
}
