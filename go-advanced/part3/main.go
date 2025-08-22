package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	done := make(chan interface{}, 10)

	cows := make(chan interface{}, 100)
	pigs := make(chan interface{}, 100)

	go func(){
		cows <- "moo"
	}()

	go func() {
		pigs <- "oink"
	}()

	wg.Add(1)
	go consumeCows(done, cows)
	wg.Add(1)
	go consumePigs(done, pigs)

}
