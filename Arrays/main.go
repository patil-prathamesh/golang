package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	var intArr [3]int32
	intArr[1] = 9
	fmt.Println(intArr[:])
	fmt.Println(&intArr[1])

	// var intArr2 = [3]int{1,2,3}
	intArr2 := []int{1, 2, 3, 7, 9, 4}
	fmt.Println(intArr2)

	var mySlice []int = []int{1, 2, 3}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 90)
	fmt.Println(mySlice)

	slice2 := []int{12, 3, 4, 5, 6}
	fmt.Println(mySlice)
	mySlice = append(mySlice, slice2...)
	fmt.Println(mySlice)

	var map1 map[string]int = make(map[string]int)
	var map2 = map[string]int{}
	map1["patil"] = 21
	var val, ok = map1["patill"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("Invalid")
	}
	map2["cat"] = 22
	map2["box"] = 23
	fmt.Println(map2)
	fmt.Println("---------------")

	arr1 := [4]int{1, 2, 3}
	slice1 := []int{4, 5, 6}
	map3 := map[int]int{}
	map3[3] = 99
	fmt.Println(arr1, slice1, map3)
	delete(map3, 3)
	map3[1] = 100
	map3[2] = 200
	map3[3] = 300
	fmt.Println(arr1, slice1, map3)
	for i, v := range arr1 {
		fmt.Println(i, v)
	}
	for _, v := range slice1 {
		fmt.Println(v, " **")
	}
	for k, v := range map3 {
		fmt.Println(k, v, " &&")
	}

	n := 1000000
	var ts = []int{}
	var ts1 = make([]int, 0, n)
	fmt.Println(timeLoop(ts, n))
	fmt.Println(timeLoop(ts1, n))
	fmt.Println(time.Now())

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
