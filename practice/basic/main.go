package main

import (
	"fmt"
	"math"
)

func main() {
	one()
	two()
	three()
	fmt.Println(reverse(343))
	fmt.Println(palindrome(343))
	fmt.Println(armstrong(2000))
	getfibbonacci(5)
}
func getfibbonacci(n int) {
	a := 0
	b := 1
	fmt.Printf("the fibonacci numbers are %v %v",a,b)
	for i:=0; i < n-2; i++ {
		c := a+b
		a = b
		b = c
		fmt.Printf(" %v",b)
	}
}
func armstrong(num int) bool {
	duplicate := num
	temp := num
	var digits int
	for temp > 0 {
		temp /= 10
		digits++
	}
	var sum float64
	for num > 0 {
		rem := num % 10
		sum += math.Pow(float64(rem),float64(digits))
		num /= 10
	}
	fmt.Println(sum, "***")
	return float64(duplicate) == sum
}
func palindrome(num int) bool {
	duplicate := num
	reverse := 0
	for num > 0 {
		rem := num % 10
		reverse = reverse * 10 + rem
		num /= 10
	}
	return duplicate == reverse
}
func reverse(num int) int {
	//345
	ans := 0
	for num > 0 {
		rem := num % 10
		ans = ans * 10 + rem
		num /= 10
	}
	return ans
}
func three() {
	num := 3458
	var ans int
	var rem int
	for num > 0 {
		rem = num % 10
		ans += rem
		num /= 10
	}
	fmt.Println(ans, "--")

}
func two() {
	a, b, c := 19,82,3

	if a > b && a > c {
		fmt.Println("A is biggest")
	}else if b > a && b > c {
		fmt.Println("B is biggest")
	}else {
		fmt.Println("C is biggest")
	}
}

func one() {
	var num int = 45
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
