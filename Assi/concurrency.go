package main

import (
	"fmt"
	"time"
)

func getValue(val int, channel chan int) {
	for i := 0; i < val; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(i)

	}
	fmt.Println(<-channel)

}

func main() {

	fmt.Println("Concurrency")
	channel := make(chan int)
	go getValue(5, channel)
	go getValue(5, channel)
	channel <- 34
	fmt.Println(channel)

}
