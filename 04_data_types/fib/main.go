package main

import "fmt"

func fibonacci(n int) []int {
	nums := make([]int, n)

	if n < 2 {
		return nums
	}

	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums
}

func main() {
	var num int

	fmt.Print("Fibonacci sequence?")
	fmt.Scanln(&num)
	fmt.Println("Fibonacci sequence:", fibonacci(num))
}
