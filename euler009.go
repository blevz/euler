package main

import "fmt"

func main() {
	for a := 0; a < (1000/3)+1; a++ {
		for b := a + 1; b < 1000-a; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println("a: ", a, " b: ", b, " c: ", c)
				fmt.Println(a * b * c)
			}
		}
	}

}
