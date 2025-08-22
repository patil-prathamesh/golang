package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// func child1(ctx context.Context) {
// 	for {
// 		select{
// 		case <-ctx.Done():
// 			return
// 		default:
// 			fmt.Println("Hello")
// 		}
// 	}
// }

func fun1(ctx context.Context) {
	defer func() {
		fmt.Println("bye bye")
		wg.Done()
	}()
	for{
		select{
		case <-ctx.Done():
			return
		default:
			fmt.Println(time.Now())
		}
	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.TODO(),time.Now().Add(time.Second*5))
	defer cancel()
	wg.Add(1)
	go fun1(ctx)
	wg.Wait()
}

/* 
withCancel()

WithDeadline(ctx,deadline) current time 3 PM deadline is 3:10 PM
Withtimeout(ctx,time.Second*3) maybe after 10,20 seconds

both can be cancel using cancel func or deadline or timeout
*/