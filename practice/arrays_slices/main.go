package main 

import "fmt"

func main() {
	var nums []int = []int{11,42,33,994,5,42,42,42,11,11}
	fmt.Println(sum(nums))
	fmt.Println(largest(nums)," largest")
	frequency(nums)
}

func frequency(nums []int) {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	fmt.Println(m)
}

func largest(nums []int) int {
	biggest := nums[0]
	n := len(nums)
	for i:=0; i< n; i++ {
		if nums[i] > biggest{
			biggest = nums[i]
		}
	}
	return biggest
}

func sum(nums []int) int {
	total := 0
	for _,v := range nums {
		total += v
	}
	return total
}