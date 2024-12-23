package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func givemeanumber() int {
	return -1
}

func main() {
	// if statements
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// fmt.Println(num) => error!

	// Switch statement
	// sec := time.Now().Unix()
	// rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	default:
		fmt.Print("no match...")
	}

	// Switch multi values
	region, continent := location("Irvine")
	fmt.Printf("\n\nJohn works in %s, %s\n", region, continent)

	// Call function from switch case
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}
	fmt.Println(time.Now().Weekday().String())

	// omit condition
	r := rand.Float64()
	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}

	// Make the logic fall through to the next case
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}

	// Looping
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("\n\nsum of 1..100 is", sum)

	// Empty prestatements and poststatements
	var num int64
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}

	// Infinite loops and break
	sum = 10
	for num := 1; num <= 100; num++ {
		if num == sum {
			break
		}

	}

	// Continue
	sum = 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)

	// Defer
	for i := 0; i < 5; i++ {
		defer fmt.Println("Deferred:", i)
		fmt.Println("Regular:", i)
	}

	// close file
	newfile, err := os.Create("learnGo.txt")
	if err != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	defer newfile.Close()

	if _, err = io.WriteString(newfile, "Learning Go!"); err != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}
	newfile.Sync()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")

	// Revover
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}
