package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"math"
)

func main() {
	pastDistinct := make([]int, 0)
	for x := 10; x < int(math.MaxInt32); x++ {

		if len(common.PrimeFactorizationMap(x)) == 4 {
			fmt.Println(x)
			pastDistinct = append(pastDistinct, x)
		} else {
			pastDistinct = make([]int, 0)
		}
		if len(pastDistinct) == 4 {
			fmt.Println(pastDistinct)
			break
		}
	}
}
