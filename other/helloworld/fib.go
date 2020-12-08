package main

func Fib(a int) int {
	if a == 1 || a == 2 {
		return 1
	} else if a > 2 {
		return Fib(a-1) + Fib(a-2)
	} else {
		return 0
	}
}
