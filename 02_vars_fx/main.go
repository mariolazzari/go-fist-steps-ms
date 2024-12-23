package main

import (
	"fmt"
	"os"
)

func main() {
	// command line args
	fmt.Println("Total command line args:", len(os.Args))
	for idx, arg := range os.Args {
		fmt.Printf("Index %v - %q\n", idx, arg)
	}
}
