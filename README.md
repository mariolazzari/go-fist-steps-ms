# First steps in Go

[Link](https://learn.microsoft.com/it-it/training/paths/go-first-steps/)

## Intro

### Install

```go
go version
```

### Go playground

[Link](https://go.dev/play/)

### First program

Official documentation for [fmt](https://pkg.go.dev/fmt) package.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

```sh
go run main.go
```

[fmt](https://pkg.go.dev/fmt) package

## Variables and functions

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// variables
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)

    // constants
    const (
        StatusOK              = 0
        StatusConnectionReset = 1
        StatusOtherError      = 2
    )

	// command line args
	fmt.Println("Total command line args:", len(os.Args))
	for idx, arg := range os.Args {
		fmt.Printf("Index %v - %q\n", idx, arg)
	}
}
```

## Control flow

### if

### switch

### for

### Defer

In Go, a defer statement postpones the running of a function (including any parameters) until the function that contains the defer statement finishes. Generally, you defer a function when you want to avoid forgetting about tasks like closing a file or running a cleanup process.

### Panic function

The built-in panic() function stops the normal flow of control in a Go program. When you use a panic call, any deferred function calls run normally. 

### Recover function

Sometimes you might want to avoid a program crash and instead report the error internally. Or perhaps you want to clean up the mess before letting the program crash.
