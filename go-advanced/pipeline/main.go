package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _,v := range nums {
			fmt.Println("Wrote...")
			out <- v
		}
		close(out)	
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			fmt.Println("Read....")
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	nums := []int{2,3,4,7,1}

	dataChannel := sliceToChannel(nums)

	finalChannel := sq(dataChannel)

	for v := range finalChannel {
		fmt.Println(v)
	}
}