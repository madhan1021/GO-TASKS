// to increment and decrement both at a same time using sync.mutex
package main

import (
	"fmt"
	"sync"
)

const count = 1000

type Counter struct {
	Count      int
	sync.Mutex //to store the lock and unlock function automatically within the struct
}

func main() {
	counter := Counter{Count: 0}
	done := make(chan struct{}) //to show the function is done so we are passing a empty struct to the channel to notify
	// goroutine to increment the num value
	go func() {
		for i := 0; i < count; i++ {
			counter.Lock() //it locks the gorotuine until it finishes
			counter.Count++
			counter.Unlock() //it unlocks once the job is done so that next gorotuine can perform.
		}
		done <- struct{}{}
	}()
	//goroutine to decrement the num value
	go func() {
		for i := 0; i < count; i++ {
			counter.Lock()
			counter.Count--
			counter.Unlock()
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(counter.Count)
}
