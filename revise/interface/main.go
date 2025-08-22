package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r *Rectangle) Area() float64 {
	return r.Breadth * r.Length
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Length * r.Breadth)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64{
	return 2 * math.Pi * c.Radius
}

func main() {
	rectangle := &Rectangle{10,14}
	circle := &Circle{4}
	shapes := []Shape{rectangle, circle}

	for _, v := range shapes {
		fmt.Println(v.Area(), " ", v.Perimeter())
	}
}
