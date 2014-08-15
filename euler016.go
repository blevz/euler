package main

import (
	"fmt"
	"math/big"
)

func sumOfDigits(num big.Int) int {
	word := num.String()
	sum := 0
	for _, ch := range word {
		sum += int(ch - '0')
	}
	return sum
}

func main() {
	product := big.NewInt(1)
	for x := 0; x < 1000; x++ {
		product.Mul(product, big.NewInt(2))
	}
	fmt.Println("2^1000:  ", product.String(), " sum is: ", sumOfDigits(*product))
}
