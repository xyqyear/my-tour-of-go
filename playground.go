package main

import (
	"fmt"
)

func fib(a int) int {
	if a == 1 || a == 2 {
		return 1
	} else if a > 2 {
		return fib(a-1) + fib(a-2)
	} else {
		return 0
	}
}

func main() {
	fmt.Println(fib(7))
}
