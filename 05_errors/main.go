package main

import (
	"fmt"
	"log"
	"os"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetPrefix("main(): ")
	log.SetOutput(file)
	employee, err := getInformation(1001)
	if err != nil {
		// Something is wrong. Do something.
		log.Fatal("Something is wrong", err)
		log.Fatal(err)
	} else {
		fmt.Print(employee)
	}
}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(1000)
	return employee, err
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
