package main

import (
	"errors"
	"fmt"
	// "unsafe"
)

func printMe(value string) {
	fmt.Print("Hello from func", value)
}

func division(num1 int, num2 int) (int, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("cannont divide by zero")
		return 0, 0, err
	}
	var result = num1 / num2
	var reminder = num1 % num2
	return result, reminder, nil
}

func main() {
	Greet()

	// var x int8
	// var xx float32
	// var name string
	// name := "patil"
	// var myBoolean bool = true
	// name1 := 12
	// var name string = "patil"
	// var name = "patil"
	// name := "patil"
	// num1, num2, num3 := 1, 2, 3
	// const myConst = "p"
	// fmt.Print(num1, num2, num3)
	// printMe("hello")
	// val1, val2, err := division(8, 0)
	// if err != nil {
	// 	fmt.Print(err)

	// } else {
	// 	fmt.Printf("Your value is %v and %v", val1, val2)
	// }

	// name, err := greet("patil")
	// fmt.Print(name, " ", err)
	// value1 := 9
	// switch value1 {
	// case 1:
	// 	fmt.Print("Hello")
	// case 2:
	// 	fmt.Print("helooooo")
	// case 3, 4, 5, 6, 7:
	// 	fmt.Print("byee")
	// default:
	// 	fmt.Println("oho")
	// }
}

func greet(name string) (string, error) {
	if name == "patil" {
		return ":(", errors.New("name matched, please change name")
	}
	return name, nil
}
