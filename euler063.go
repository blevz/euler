package main

import (
	"fmt"
	_ "github.com/blevz/euler/common"
	"math"
	"math/big"
)

//Max will always be 10 because

func main() {
	num := 0
	//x represents the number of digits
	for x := 1; x < 100; x++ {
		low := math.Pow(float64(10), float64(x-1))
		high := math.Pow(float64(10), float64(x))
		min := math.Pow(low, float64(1)/float64(x))
		max := math.Pow(high, float64(1)/float64(x))
		for a := int(min); a < int(max)+1; a++ {
			base := big.NewInt(int64(a))
			exp := big.NewInt(int64(x))
			base.Exp(base, exp, nil)
			if len(base.String()) == x {
				fmt.Println(base.String())
				num++
			}
		}
		if min > 9 {
			break
		}

	}
	fmt.Print(num)

}
