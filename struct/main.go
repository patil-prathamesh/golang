package main

import (
	"fmt"
)

type gasEngine struct {
	mpg uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

type engine interface {
	milesLeft() uint8
}

func (e *gasEngine) milesLeft() (uint8) {
	return e.gallons*e.mpg
}

func (e *electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there")
	}else {
		fmt.Println("Need to fuel up first")
	}
}

// type owner struct {
// 	name string
// 	address string
// }

func main() {
	// myEngine := gasEngine{mpg: 45, gallons: 23,ownerInfo: owner{name: "psp"}}
	// myEngine := gasEngine{23,45,owner{"patil"}}

	// myEngine := gasEngine{23,45,owner{"prthmsh","kalamboli"},12}
	// fmt.Println(myEngine.address)
	// fmt.Println(myEngine.int)

	myStruct := struct{
		age int
		name string
	}{12,"patil"}
	fmt.Println(myStruct)

	myEngine := gasEngine{25, 15}
	canMakeIt(&myEngine, 50)
	engine2 := electricEngine{25, 15}
	canMakeIt(&engine2, 50)
}