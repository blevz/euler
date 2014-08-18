package main

import (
	"fmt"
	"math/big"
)

func main() {
	product := big.NewInt(1)
	for x := 1; x < 100; x++ {
		product.Mul(product, big.NewInt(int64(x)))
	}
	sum := 0
	for _, v := range product.String() {
		sum += int(v - '0')
	}
	fmt.Println(sum)
}
