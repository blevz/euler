package main

import (
	"fmt"
	"math"
)

var (
	MappedVals map[int64]struct{}
	ArrOfVals  []int64
)

func GeneratePentagonalNumbers(c chan<- int64) {
	for n := int64(1); true; n++ {
		c <- ((n * ((3 * n) - 1)) / 2)
	}
}

func main() {
	CurMin := int64(math.MaxInt64)
	MappedVals = make(map[int64]struct{}, 0)
	ArrOfVals = make([]int64, 0)
	PentagonalNumChan := make(chan int64)
	go GeneratePentagonalNumbers(PentagonalNumChan)

	x := <-PentagonalNumChan
	MappedVals[x] = struct{}{}
	ArrOfVals = append(ArrOfVals, x)

	for ph := range PentagonalNumChan {
		if ph-ArrOfVals[len(ArrOfVals)-1] > CurMin {
			fmt.Println(CurMin)
			return
		}
		for _, pj := range ArrOfVals {
			pk := ph - pj
			if _, exists := MappedVals[pk]; exists {
				d := pk - pj
				if d < 0 {
					d *= -1
				}
				if _, exists := MappedVals[d]; exists {
					CurMin = d
					fmt.Println(CurMin)
				}
			}
		}
		//fmt.Println(len(ArrOfVals))

		MappedVals[ph] = struct{}{}
		ArrOfVals = append(ArrOfVals, ph)
	}
}
