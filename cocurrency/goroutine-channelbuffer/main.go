package main

import (
	"fmt"
	"time"
)

func Sender(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("sent value:", i)
	}
	close(c)
}

func Receiver(c chan int) {
	for {
		if value, ok := <-c; ok {
			fmt.Println("recved value:", value)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
}

func main() {
	channel := make(chan int, 2)
	go Sender(channel)
	Receiver(channel)
}
