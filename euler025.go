package main

import (
	"fmt"
	"math/big"
)

//Use big int

func main() {
	num := 2
	a := big.NewInt(1)
	b := big.NewInt(1)

	for len(b.String()) < 1000 {
		temp := big.NewInt(0)
		temp.Add(a, b)
		a = b
		b = temp
		num++
	}
	fmt.Println(num, b)

}
