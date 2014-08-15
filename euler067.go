package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func main() {
	tree := common.FileToTriangleArray("euler067_data.txt")
	for h := tree.GetDepth() - 2; h >= 0; h-- {
		for x := 0; x < h+1; x++ {
			val := common.Max_int2(tree.GetVal(h+1, x).(int), tree.GetVal(h+1, x+1).(int))
			val += tree.GetVal(h, x).(int)
			tree.SetVal(h, x, val)
		}
	}
	//Start at the largest depth
	fmt.Println(tree.GetVal(0, 0))
}
