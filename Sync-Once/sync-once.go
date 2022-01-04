//Lazy loading
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var config map[string]string
var configOnce sync.Once = sync.Once{} //will not allow other goroutines to perform if one time it gets loaded

func loadConfig() {
	time.Sleep(100 * time.Millisecond)
	log.Println("Loading Configuration") //to test the loading of config and we need only 1 time to load.
	config = map[string]string{
		"hostname": "localhost",
		"port":     "8080",
	}
}

//lazy loading basic structure
func getConfig(key string) string {
	if config == nil {
		configOnce.Do(loadConfig) //Only accepts the first function call and the rest will get ignored
	}
	return config[key]
}
func doSomething(done chan struct{}) {
	tf := getConfig("hostname")
	tf1 := getConfig("port")
	fmt.Println(tf)
	fmt.Println(tf1)
	done <- struct{}{}
}
func main() {
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		go doSomething(done)
	}
	for i := 0; i < 5; i++ {
		<-done
	}
}
