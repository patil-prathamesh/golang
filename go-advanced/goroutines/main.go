// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func someFunc(num string, done chan bool) {
// //     fmt.Println(num)
// // 	time.Sleep(4*time.Second)
// //     done <- true // Signal completion
// // }

// // func secondFunc(num string, done chan bool) {
// //     fmt.Println(num)
// // 	time.Sleep(2*time.Second)
// // 	done <- true
// // }

// // func thirdFunc(num string, done chan bool) {
// // 	fmt.Println(num)
// // 	time.Sleep(3*time.Second)
// // 	done <- true
// // }

// // chan bool      // Bidirectional (can send and receive)
// // <-chan bool    // Receive-only (can only receive)
// // chan<- bool    // Send-only (can only send)
// //receive only or read only
// func doWork(done <-chan bool) {
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		default:
// 			fmt.Println("hello")
// 		}
// 	}
// }

// func main() {
// 	// done1 := make(chan bool)
// 	// done2 := make(chan bool)
// 	// done3 := make(chan bool)

// 	// go someFunc("2", done1)
// 	// go thirdFunc("5", done2)
// 	// go secondFunc("9", done3)
// 	// fmt.Println("hello")

// 	// select {
// 	// case chan1 := <- done1:
// 	// 	fmt.Println(chan1)
// 	// case chan2 := <- done2:
// 	// 	fmt.Println(chan2)
// 	// case chan3 := <- done3:
// 	// 	for i := range 5 {
// 	// 		fmt.Print(i, " ")
// 	// 	}
// 	// 	fmt.Println(chan3)
// 	// }
// 	// fmt.Println("byee")

// 	done := make(chan bool)
// 	go doWork(done)
// 	time.Sleep(time.Second*3)
// 	close(done)
// }

// package main

// import (
//     "fmt"
//     "time"
// )

// // Send-only channel - can only send data
// func worker(id int, jobs chan<- string) {
//     for i := 1; i <= 3; i++ {
//         job := fmt.Sprintf("Worker %d - Job %d", id, i)
//         jobs <- job
// 		fmt.Println("did work ", id)         // Send job to channel
//         time.Sleep(time.Millisecond * 500)
//     }
// }

// // Receive-only channel - can only receive data
// func collector(results <-chan string) {
//     for result := range results {
//         fmt.Println("Collected:", result)
//     }
// }

// func main() {
//     jobs := make(chan string, 10)      // Bidirectional channel

//     // Start workers (send-only access)
//     go worker(1, jobs)
//     go worker(2, jobs)

//     // Start collector in background
//     go collector(jobs)

//     time.Sleep(time.Second * 4)
//     close(jobs)
//     time.Sleep(time.Millisecond * 100) // Let collector finish
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// //send only receive only

// func worker(id int, jobs chan<-string) {
// 	for i := 1; i <= 2; i++ {
// 		msg := fmt.Sprintf("worker %d job %d",id, i)
// 		jobs <- msg
// 		time.Sleep(time.Second*1)
// 	}
// }

// func collector(jobs <-chan string) {
// 	for result := range jobs {
// 		fmt.Println("collected: ", result)
// 	}
// }

// func main() {
// 	jobs := make(chan string, 10)

// 	go worker(1,jobs)

// 	go collector(jobs)

// 	time.Sleep(6 * time.Second)  // Wait for worker to finish
// 	close(jobs)
// 	time.Sleep(100 * time.Millisecond)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func doWork(done1 <-chan bool) {
// 	for{
// 		select{
// 		case <-done1:
// 			return
// 		default:
// 			fmt.Println("hello")
// 		}
// 	}
// }

// func sum(arr []int, result chan int) {
// 	var sum int
// 	for _, v := range arr {
// 		sum += v
// 	}
// 	result <- sum
// }

// type object struct {
// 	age int
// 	name string
// }

// func obj(result chan object) {
// 	time.Sleep(2*time.Second)
// 	result <- object{23,"nikhil"}
// }

// func main() {
// 	arr := []int{1,2,3,4,5,6}
// 	result1 := make(chan object)
// 	result := make(chan int)
// 	go sum(arr, result)
// 	go obj(result1)
// 	fmt.Println(<-result)
// 	fmt.Println(<-result1)
// 	// done := make(chan bool)

// 	// go doWork(done)

// 	// time.Sleep(3*time.Second)
// 	// close(done)
// }

package main

import (
	"fmt"
)

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _,v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func(){
		for v := range in {
			out <- v*v
		}
		close(out)
	}()
	return out
}

func main() {
	nums := []int{1,2,3,4,5}

	//stage 1
	dataChannel := sliceToChannel(nums)
	fmt.Println("bye")
	//stage 2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel{
		fmt.Println(n)
	}
	fmt.Print("hello")
}