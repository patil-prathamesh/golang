package main

import (
	"fmt"
	"strings"
)

func main() {
	two()
}

type Student struct {
	Name string
	Grade string
}

func three() {}

func two() {
	students := []Student{
        {"Alice", "A"},
        {"Bob", "B"},
        {"Charlie", "A"},
        {"Diana", "C"},
        {"Eve", "B"},
        {"Frank", "A"},
        {"Grace", "C"},
    }
	result := groupByGrade(students)
	for k, v:= range result {
		fmt.Printf("Grade %v: ", k)
		for _, v := range v {
			fmt.Print(v.Name, ", ")
		}
		fmt.Println("")
	}
	// fmt.Println(result)
}

func groupByGrade(students []Student) map[string][]Student {
	group := map[string][]Student{}
	for _, v := range students {
		group[v.Grade] = append(group[v.Grade], v)
	}
	return group
}

func one() {
	str := "Hello world hello"
	words := strings.Split(str, " ")
	occurrences := map[string]int{}
	for _, v := range words {
		lowerWord := strings.ToLower(v)
		occurrences[lowerWord]++
	}
	fmt.Println(occurrences)
}
