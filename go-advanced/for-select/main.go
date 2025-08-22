package main

import (
	"fmt"
	"time"
)

// make communication async beween goroutines buffered channel

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Do work")
		}
	}
}
func main() {
	done := make(chan bool)
	go doWork(done)

	time.Sleep(time.Second * 3)
	close(done)
}