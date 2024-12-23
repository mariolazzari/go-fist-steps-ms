package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {
	// Arrays
	var a [3]int
	fmt.Println(a)
	a[1] = 1
	fmt.Println(a)
	a[len(a)-1] = 2
	fmt.Println(a)

	// arrays init
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	// ellipsys
	q := [...]int{1, 2, 3}
	fmt.Println(q)

	// Slices
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))

	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))

	// append
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}

	// remove
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {
		fmt.Println("Before", letters, "Remove ", letters[remove])
		letters = append(letters[:remove], letters[remove+1:]...)
		fmt.Println("After", letters)
	}

	// copy slice
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	letters = []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	slice2 = letters[1:4]

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)

	// Maps
	studentAges := map[string]int{
		"Mario":     49,
		"Mariarosa": 48,
	}
	fmt.Println(studentAges)

	// make
	studentsAges := make(map[string]int)

	// add element
	studentsAges["mario"] = 21
	studentsAges["maria"] = 19
	fmt.Println("Mario's age is", studentsAges["mario"])

	// check if element exists
	age, exists := studentsAges["gino"]
	if exists {
		fmt.Println("Gino is", age)
	}

	// delete element
	delete(studentsAges, "mario")
	fmt.Println(studentsAges)

	// looping into a map
	for name, age := range studentAges {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// only val
	for _, age := range studentAges {
		fmt.Println(age)
	}
	// only key
	for key := range studentAges {
		fmt.Println(key)
	}

	// Structs
	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}
	employee.LastName = "Doe"
	fmt.Println(employee.FirstName)

	employees := []Employee{
		{
			Person: Person{
				LastName: "Doe", FirstName: "John",
			},
		},
		{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}
