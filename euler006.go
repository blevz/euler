package main

import "fmt"

func main() {
	sum1 := 0
	sum2 := 0
	for x := 0; x <= 100; x++ {
		sum1 += x
		sum2 += x * x
	}
	sum1 = sum1 * sum1
	fmt.Println(sum1)
	fmt.Println(sum2)
	diff := sum1 - sum2
	fmt.Println(diff)
}
