package main

import (
	"fmt"
	"strconv"
)

func factorial(num int) int {
	product := 1
	for x := 2; x <= num; x++ {
		product *= x
	}
	return product
}

//You can view the lexigraphic ordering of permutations as a tree with h! leaves

func main() {
	curNum := 1
	finalNum := 1000000
	length := 10

	used := make([]bool, 10)
	for x := 0; x < 10; x++ {
		used[x] = false
	}
	finalAnswer := ""
	for depth := 0; depth < length; depth++ {
		depthFact := factorial(length - 1 - depth)
		indexOfBranch := 0
		for i, v := range used {
			if !v {
				if depthFact*(indexOfBranch+1)+curNum > finalNum {
					finalAnswer = finalAnswer + strconv.Itoa(i)
					used[i] = true
					curNum += depthFact * (indexOfBranch)
					break
				} else {
					indexOfBranch++
				}

			}
		}
	}
	fmt.Println(finalAnswer)
}
