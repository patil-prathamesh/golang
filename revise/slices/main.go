package main

import "fmt"

func main() {
	one()
	two()
	slice := []int{1, 2, 3}
    modifySlice(slice)
    fmt.Println(slice) 
}
func modifySlice(s []int) {
    s[0] = 100
    s = append(s, 200)
}

func one() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums2 := nums[2:8]
	nums2 = append(nums2,11,12,13)
	fmt.Println(nums2)
}

func two() {
	numbers := []int{5, 10, 15, 20, 25, 30}
	numbers = append(numbers[:2], numbers[3:]...)
	numbers = append(numbers[:1], append([]int{12}, numbers[1:]...)...)
	fmt.Println(numbers)
}
