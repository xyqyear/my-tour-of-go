package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("sending value", sum)
	time.Sleep(time.Second)
	c <- sum // send sum to c
	fmt.Println("sent value", sum)
	time.Sleep(time.Second)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x := <-c
	fmt.Println("got x:", x)
	time.Sleep(time.Second)
	y := <-c
	fmt.Println("got y:", y)
	time.Sleep(time.Second)
	fmt.Println(x, y, x+y)
}
