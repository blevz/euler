package main

import (
	"fmt"
	"math"
)

func GenTriangleNums(output chan<- int) {
	GenSequence(func(a int) int {
		return a * (a + 1) / 2
	}, output)
}

func GenPentagonNums(output chan<- int) {
	GenSequence(func(a int) int {
		return a * (3*a - 1) / 2
	}, output)
}

func GenHexNums(output chan<- int) {
	GenSequence(func(a int) int {
		return a * (2*a - 1)
	}, output)
}

func GenSequence(f func(int) int, output chan<- int) {
	for x := 1; x < int(math.MaxInt32); x++ {
		output <- f(x)
	}
	close(output)
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	go GenPentagonNums(a)
	go GenHexNums(b)
	go GenTriangleNums(c)
	curPent := <-a
	curHex := <-b
	curTri := <-c
	curPent = <-a
	curHex = <-b
	curTri = <-c

	fmt.Println(curTri, curPent, curHex)
	for curPent != curTri || curTri != curHex {
		//Determine the smallest num and get the next num
		if curTri <= curPent && curTri <= curHex {
			curTri = <-c
		} else if curHex <= curPent && curHex <= curTri {
			curHex = <-b
		} else {
			curPent = <-a
		}
	}
	fmt.Println(curTri, curPent, curHex)
	curPent = <-a
	curHex = <-b
	curTri = <-c
	for curPent != curTri || curTri != curHex {
		//Determine the smallest num and get the next num
		if curTri <= curPent && curTri <= curHex {
			curTri = <-c
		} else if curHex <= curPent && curHex <= curTri {
			curHex = <-b
		} else {
			curPent = <-a
		}
	}
	fmt.Println(curTri, curPent, curHex)
}
