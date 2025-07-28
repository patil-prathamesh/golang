package utils

func Power(base, exponent int) int {
    result := 1
    for i := 0; i < exponent; i++ {
        result *= base
    }
	product := Product(4,5)
    return result+product
}