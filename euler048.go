package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"math/big"
)

func main() {
	sum := big.NewInt(0)
	p := big.NewInt(int64(common.Pow(10, 10)))
	fmt.Println(p)
	for x := 1; x <= 1000; x++ {
		b := big.NewInt(int64(x))
		for y := 1; y < x; y++ {
			//fmt.Println(y, b)
			b.Mul(b, big.NewInt(int64(x)))
			b.Mod(b, p)
		}
		//fmt.Println(b)
		sum.Add(sum, b)
	}
	sum.Mod(sum, p)
	fmt.Println(sum)
}
