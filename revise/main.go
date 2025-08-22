package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x, y int) (ans int) {
	ans = x + y
	return
}

func swap(str1 string, str2 string) (string, string) {
	return str2, str1
}

type Level uint8

const (
	one Level = iota + 1
	two
	three
	four
	five
)

var (
	pp = 22
	np = 19
	sp = 50
	up = 44
)

func main() {
	fmt.Println("hello")
	fmt.Println(rand.Intn(30))
	fmt.Println(math.Pi)
	fmt.Println(add(4, 5))

	var a, b, c = 34, "patil", false
	fmt.Println(a, b, c)

	y, z := swap("nikhil", "prthmsh")
	fmt.Println(y, z)
	fmt.Println(one, two, three, four, five, pp, np, sp, up)

	f := 3.12
	zz := uint8(f)
	fmt.Println(zz)
	Loops()
	deferr()
	pointers()
	structs()
	Slices()
}

func sone() {
	slice1 := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(len(slice1), cap(slice1), slice1[2], slice1[len(slice1)-1])
}

func stwo() {
	s := []int{10,20,30,40,50}
	s2 := []int{60,70}
	s = append(s, s2...)
	fmt.Println(s)
}
func sthree() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:7]
	s2 := arr[:3]
	s3 := arr[len(arr)-3:]
	fmt.Printf("s1 -> %v, %v /n", len(s1),cap(s1))
	fmt.Printf("s1 -> %v, %v /n", len(s2),cap(s2))
	fmt.Printf("s1 -> %v, %v /n", len(s3),cap(s3))
}

func sfive() {
	matrix := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}

	for _, v := range matrix{
		for _, v := range v {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}

	for i, v := range matrix{
		var ans int
		for _,v := range v {
			ans += v
		}
		fmt.Printf("Row %v -> %v", i, ans)
		fmt.Println("")
	}

	for i := range len(matrix[0]) {
		ans := 0
		for j := range len(matrix) {
			ans += matrix[j][i]
		}
		fmt.Println(ans, " **")
	}
}

func question8() {
    fmt.Println("\n=== Question 8 ===")
    
    slice1 := []int{1, 2, 3, 4, 5}
    
    // Shallow copy (assignment)
    slice2 := slice1
    fmt.Printf("Original slice1: %v\n", slice1)
    fmt.Printf("slice2 = slice1: %v\n", slice2)
    
    slice2[0] = 999
    fmt.Printf("After modifying slice2[0]:\n")
    fmt.Printf("slice1: %v\n", slice1)
    fmt.Printf("slice2: %v\n", slice2)
    
    // Deep copy using copy()
    slice1 = []int{1, 2, 3, 4, 5} // reset
    slice3 := make([]int, len(slice1))
    copy(slice3, slice1)
    
    fmt.Printf("\nAfter deep copy:\n")
    fmt.Printf("slice1: %v\n", slice1)
    fmt.Printf("slice3: %v\n", slice3)
    
    slice3[0] = 888
    fmt.Printf("After modifying slice3[0]:\n")
    fmt.Printf("slice1: %v\n", slice1)
    fmt.Printf("slice3: %v\n", slice3)
}

func Slices() {
	sone()
	stwo()
	sthree()
	sfive()
	question8()
}

func structs() {
	type User struct {
		ID string `json:"id"`
		UserName string `json:"username"`
		Email string `json:"email"`
		Active bool `json:"active"`
	}

	type Product struct {
		Name string `json:"name`
		Price float32 `json:"price"`
		Category string `json:"category"`
		Tags []string `json:"tags"`
		Instock bool `json:"instock"`
	}

	type Coordiantes struct {
		Latitude string
		Longitude string
	}

	type Address struct {
		Street string `json:"street"`
		City string `json:"city"`
		State string `json:"state"`
		ZipCode string `json:"zipcode"`
		Coordiantes
	}

	type UserInfo struct {
		User string `json:"user"`
		Text string `json:"Text`
	}

	type Post struct {
		PostID string `json:"postId"`
		Author string `json:"author"`
		Content string `json:"content"`
		Likes int64
		Comments []UserInfo
		TimeStamp time.Time
	}

	type UserData struct {
		ID int64 `json:"id"`
		Name string `json:"name"`
	}

	type Response struct {
		Status string `json:"status"`
		Data struct{
			Users []UserData `json:"users"`
			Total int `json:"total"`
		} `json:"data"`
		Message string `json:"message"`
	}


	
}

func pointers() {
	i, j := 12,5555
	fmt.Println(i)
	p := &i
	*p = 90909090
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)

	p = &j
	fmt.Println(*p*88)
}

func Index[T comparable] (s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func deferr() {
	defer fmt.Println("deferrr")
	defer fmt.Println("first")
	defer fmt.Println("first2")
	defer fmt.Println("first3")
	defer fmt.Println("first4")
	fmt.Println("hello")
}
