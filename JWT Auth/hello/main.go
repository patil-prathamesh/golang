package main

import "fmt"

type obj struct {
	name string
	age int
	key
}

type key struct {
	pair int
}

func main() {
	obj1 := obj{
		name: "hello",
		age: 34,
		key: key{
			pair: 34,
		},
	}

	newobj := obj{
		"nikhil",
		34,
		key{
			23,
		},
	}
	// fmt.Print("hii")
	// obj1 := obj{"patil", 12}
	// obj3 := obj{name:"nikhil", age:12}
	// obj2 := key{pair: map[string]string{"hello":"hh"}}
	// map1 := map[int]string{1:"patil"}
	// fmt.Println(obj1,map1, obj2,obj3)
}
