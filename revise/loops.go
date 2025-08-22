package main

import (
	"fmt"
	"runtime"
	"time"
)

func Loops() {
	// for i := 1; i < 5; i++ {
	// 	fmt.Print(rand.Intn(i), "->")
	// }

	// j := 5
	// for j < 10 {
	// 	fmt.Println(j*999)
	// 	j++
	// }

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os x.")
	case "linux":
		fmt.Println(os)
	default:
		fmt.Println(os)
	}

	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().YearDay())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now())
	defer fmt.Println("Loops")
}
