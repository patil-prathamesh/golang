package main

import (
	"fmt"
)

type Number interface {
	int | int32 | float32 | float64
}

func sum[T Number](numbers []T) T {
	var result T
	for _,v := range numbers {
		result += v
	}
	return result
}

func main() {
	fmt.Println(sum([]int{1,2,3,4,5}))
}