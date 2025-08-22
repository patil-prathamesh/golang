package main
import "fmt"

func activateGiftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return  amount
	}

	return debitFunc
}

func main() {
	use1 := activateGiftCard()
	// use2 := activateGiftCard()

	fmt.Println(use1(10))
	fmt.Println(use1(5))
}